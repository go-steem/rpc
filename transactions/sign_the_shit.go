package transactions

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	_ "encoding/hex"
	"errors"
	"fmt"
	"github.com/asuleymanov/golos-go/rfc6979"
	secp256k1 "github.com/btcsuite/btcd/btcec"
	"github.com/tendermint/go-crypto"
	"log"
	"math/big"
)

func (tx *SignedTransaction) Sign_Single(priv_b []byte, data []byte) []byte {
	privKeyBytes := [32]byte{}
	copy(privKeyBytes[:], priv_b)

	////////////
	privKey := crypto.PrivKeySecp256k1(privKeyBytes)
	priv__, _ := secp256k1.PrivKeyFromBytes(secp256k1.S256(), privKey[:])
	pri_ecdsa := priv__.ToECDSA()
	sigBytes := signBuffer(data, pri_ecdsa)

	return sigBytes
}

func signBuffer(buf []byte, private_key *ecdsa.PrivateKey) []byte {
	//Debug info
	//log.Println("signBuffer buf=", hex.EncodeToString(buf))
	// Hash a message.
	alg := sha256.New()
	alg.Write(buf)

	_hash := alg.Sum(nil)

	return signBufferSha256(_hash, private_key)
}

func signBufferSha256(buf_sha256 []byte, private_key *ecdsa.PrivateKey) []byte { // *secp256k1.Signature
	//Debug info
	//log.Println("signBufferSha256 buf_sha256=", hex.EncodeToString(buf_sha256))

	var buf_sha256_clone = make([]byte, len(buf_sha256))
	copy(buf_sha256_clone, buf_sha256)

	nonce := 0

	for {
		//Debug info
		//log.Println("before call SignECDSA", "msg_sha=", hex.EncodeToString(buf_sha256_clone), "nonce=", nonce) // "msg=", hex.EncodeToString(msg),
		r, s, err := rfc6979.SignECDSA(private_key, buf_sha256_clone, sha256.New, nonce)
		//nonce = nonce.Add(nonce, big.NewInt(1))
		nonce++
		if err != nil {
			log.Println(err)
			return nil
		}

		ecsignature := &secp256k1.Signature{R: r, S: s}

		der := ecsignature.Serialize()
		lenR := der[3]
		lenS := der[5+lenR]
		//log.Println("lenR=", lenR, "lenS", lenS)

		if lenR == 32 && lenS == 32 {
			//////////////////////////////////////
			// bitcoind checks the bit length of R and S here. The ecdsa signature
			// algorithm returns R and S mod N therefore they will be the bitsize of
			// the curve, and thus correctly sized.
			key := (*secp256k1.PrivateKey)(private_key)
			curve := secp256k1.S256()
			max_counter := 4 //max_counter := (curve.H+1)*2
			for i := 0; i < max_counter; i++ {
				//for i := 0; i < (curve.H+1)*2; i++ {
				//for i := 0; ;i++ {
				pk, err := recoverKeyFromSignature(curve, ecsignature, buf_sha256_clone, i, true)

				if err == nil && pk.X.Cmp(key.X) == 0 && pk.Y.Cmp(key.Y) == 0 {
					//result := make([]byte, 1, 2*curve.byteSize+1)
					byteSize := curve.BitSize / 8
					result := make([]byte, 1, 2*byteSize+1)
					result[0] = 27 + byte(i)
					if true { // isCompressedKey
						result[0] += 4
					}
					// Not sure this needs rounding but safer to do so.
					curvelen := (curve.BitSize + 7) / 8

					// Pad R and S to curvelen if needed.
					bytelen := (ecsignature.R.BitLen() + 7) / 8
					if bytelen < curvelen {
						result = append(result, make([]byte, curvelen-bytelen)...)
					}
					result = append(result, ecsignature.R.Bytes()...)

					bytelen = (ecsignature.S.BitLen() + 7) / 8
					if bytelen < curvelen {
						result = append(result, make([]byte, curvelen-bytelen)...)
					}
					result = append(result, ecsignature.S.Bytes()...)

					//return result, nil
					return result
					//break
				}

				//log.Println(i)
			}

			//return ecsignature
		}
	}
}

func recoverKeyFromSignature(curve *secp256k1.KoblitzCurve, sig *secp256k1.Signature, msg []byte, iter int, doChecks bool) (*secp256k1.PublicKey, error) {
	// 1.1 x = (n * i) + r
	Rx := new(big.Int).Mul(curve.Params().N,
		new(big.Int).SetInt64(int64(iter/2)))
	Rx.Add(Rx, sig.R)
	if Rx.Cmp(curve.Params().P) != -1 {
		return nil, errors.New("calculated Rx is larger than curve P")
	}

	// convert 02<Rx> to point R. (step 1.2 and 1.3). If we are on an odd
	// iteration then 1.6 will be done with -R, so we calculate the other
	// term when uncompressing the point.
	Ry, err := decompressPoint(curve, Rx, iter%2 == 1)
	if err != nil {
		return nil, err
	}

	// 1.4 Check n*R is point at infinity
	if doChecks {
		nRx, nRy := curve.ScalarMult(Rx, Ry, curve.Params().N.Bytes())
		if nRx.Sign() != 0 || nRy.Sign() != 0 {
			return nil, errors.New("n*R does not equal the point at infinity")
		}
	}

	// 1.5 calculate e from message using the same algorithm as ecdsa
	// signature calculation.
	e := hashToInt(msg, curve)

	// Step 1.6.1:
	// We calculate the two terms sR and eG separately multiplied by the
	// inverse of r (from the signature). We then add them to calculate
	// Q = r^-1(sR-eG)
	invr := new(big.Int).ModInverse(sig.R, curve.Params().N)

	// first term.
	invrS := new(big.Int).Mul(invr, sig.S)
	invrS.Mod(invrS, curve.Params().N)
	sRx, sRy := curve.ScalarMult(Rx, Ry, invrS.Bytes())

	// second term.
	e.Neg(e)
	e.Mod(e, curve.Params().N)
	e.Mul(e, invr)
	e.Mod(e, curve.Params().N)
	minuseGx, minuseGy := curve.ScalarBaseMult(e.Bytes())

	// TODO: this would be faster if we did a mult and add in one
	// step to prevent the jacobian conversion back and forth.
	Qx, Qy := curve.Add(sRx, sRy, minuseGx, minuseGy)

	return &secp256k1.PublicKey{
		Curve: curve,
		X:     Qx,
		Y:     Qy,
	}, nil
}

func decompressPoint(curve *secp256k1.KoblitzCurve, x *big.Int, ybit bool) (*big.Int, error) {
	// TODO: This will probably only work for secp256k1 due to
	// optimizations.

	// Y = +-sqrt(x^3 + B)
	x3 := new(big.Int).Mul(x, x)
	x3.Mul(x3, x)
	x3.Add(x3, curve.Params().B)

	// now calculate sqrt mod p of x2 + B
	// This code used to do a full sqrt based on tonelli/shanks,
	// but this was replaced by the algorithms referenced in
	// https://bitcointalk.org/index.php?topic=162805.msg1712294#msg1712294
	y := new(big.Int).Exp(x3, curve.QPlus1Div4(), curve.Params().P)

	if ybit != isOdd(y) {
		y.Sub(curve.Params().P, y)
	}
	if ybit != isOdd(y) {
		return nil, fmt.Errorf("ybit doesn't match oddness")
	}
	return y, nil
}

func isOdd(a *big.Int) bool {
	return a.Bit(0) == 1
}

func hashToInt(hash []byte, c elliptic.Curve) *big.Int {
	orderBits := c.Params().N.BitLen()
	orderBytes := (orderBits + 7) / 8
	if len(hash) > orderBytes {
		hash = hash[:orderBytes]
	}

	ret := new(big.Int).SetBytes(hash)
	excess := len(hash)*8 - orderBits
	if excess > 0 {
		ret.Rsh(ret, uint(excess))
	}
	return ret
}

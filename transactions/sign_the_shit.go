package transactions

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	//"encoding/hex"
	//"log"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/pkg/errors"
)

//SignSingle signature of the transaction by one of the keys
func (tx *SignedTransaction) SignSingle(privB, data []byte) ([]byte, error) {
	privKeyBytes := [32]byte{}
	copy(privKeyBytes[:], privB)

	priv, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes[:])
	priEcdsa := priv.ToECDSA()

	return signBuffer(data, priEcdsa)
}

func signBuffer(buf []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	//log.Println("signBuffer buf=", hex.EncodeToString(buf))

	// Hash a message.
	alg := sha256.New()
	_, errAlg := alg.Write(buf)
	if errAlg != nil {
		return []byte{}, errAlg
	}

	_hash := alg.Sum(nil)

	return signBufferSha256(_hash, privateKey)
}

func signBufferSha256(bufSha256 []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	var bufSha256Clone = make([]byte, len(bufSha256))
	copy(bufSha256Clone, bufSha256)

	key := (*btcec.PrivateKey)(privateKey)

	for {
		ecsignature, err := key.Sign(bufSha256Clone)

		if err != nil {
			return nil, errors.Wrapf(err, "SignSingle[signBufferSha256]: ")
		}

		der := ecsignature.Serialize()
		lenR := der[3]
		lenS := der[5+lenR]
		//log.Println("lenR=", lenR, "lenS", lenS)

		if lenR == 32 && lenS == 32 {
			//////////////////////////////////////
			// bitcoind checks the bit length of R and S here. The ecdsa signature
			// algorithm returns R and S mod N therefore they will be the bitsize of
			// the curve, and thus correctly sized.
			curve := btcec.S256()
			maxCounter := 4
			for i := 0; i < maxCounter; i++ {
				pk, err := recoverKeyFromSignature(curve, ecsignature, bufSha256Clone, i, true)

				if err == nil && pk.X.Cmp(key.X) == 0 && pk.Y.Cmp(key.Y) == 0 {

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

					return result, nil
				} else if err != nil {
					return nil, err
				}
			}
		}
	}
}

func recoverKeyFromSignature(curve *btcec.KoblitzCurve, sig *btcec.Signature, msg []byte, iter int, doChecks bool) (*btcec.PublicKey, error) {
	rx := new(big.Int).Mul(curve.Params().N,
		new(big.Int).SetInt64(int64(iter/2)))
	rx.Add(rx, sig.R)
	if rx.Cmp(curve.Params().P) != -1 {
		return nil, errors.New("SignSingle[recoverKeyFromSignature]: calculated Rx is larger than curve P")
	}

	// convert 02<Rx> to point R. (step 1.2 and 1.3). If we are on an odd
	// iteration then 1.6 will be done with -R, so we calculate the other
	// term when uncompressing the point.
	ry, err := decompressPoint(curve, rx, iter%2 == 1)
	if err != nil {
		return nil, err
	}

	// 1.4 Check n*R is point at infinity
	if doChecks {
		nRx, nRy := curve.ScalarMult(rx, ry, curve.Params().N.Bytes())
		if nRx.Sign() != 0 || nRy.Sign() != 0 {
			return nil, errors.New("SignSingle[recoverKeyFromSignature]: n*R does not equal the point at infinity")
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
	sRx, sRy := curve.ScalarMult(rx, ry, invrS.Bytes())

	// second term.
	e.Neg(e)
	e.Mod(e, curve.Params().N)
	e.Mul(e, invr)
	e.Mod(e, curve.Params().N)
	minuseGx, minuseGy := curve.ScalarBaseMult(e.Bytes())

	// TODO: this would be faster if we did a mult and add in one
	// step to prevent the jacobian conversion back and forth.
	qx, qy := curve.Add(sRx, sRy, minuseGx, minuseGy)

	return &btcec.PublicKey{
		Curve: curve,
		X:     qx,
		Y:     qy,
	}, nil
}

func decompressPoint(curve *btcec.KoblitzCurve, x *big.Int, ybit bool) (*big.Int, error) {
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
		return nil, errors.New("SignSingle[decompressPoint]: ybit doesn't match oddness")
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

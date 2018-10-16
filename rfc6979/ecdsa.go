package rfc6979

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"hash"
	"math/big"
	//"log"
	//"encoding/hex"
)

// SignECDSA signs an arbitrary length hash (which should be the result of
// hashing a larger message) using the private key, priv. It returns the
// signature as a pair of integers.
//
// Note that FIPS 186-3 section 4.6 specifies that the hash should be truncated
// to the byte-length of the subgroup. This function does not perform that
// truncation itself.
func SignECDSA(priv *ecdsa.PrivateKey, hash []byte, alg func() hash.Hash, nonce int) (r, s *big.Int, err error) {
	c := priv.PublicKey.Curve
	n := c.Params().N

	//log.Println("e=", hex.EncodeToString(hash)) 		 // OK
	//log.Println("N=", hex.EncodeToString(N.Bytes()))	 // OK

	var hashClone = make([]byte, len(hash))
	copy(hashClone, hash)

	//log.Println("generateSecret- nonce=", nonce)
	generateSecret(priv /* N, priv.D, */, alg, hashClone, func(k *big.Int) bool {
		inv := new(big.Int).ModInverse(k, n)
		r, _ = priv.Curve.ScalarBaseMult(k.Bytes())
		r.Mod(r, n)

		if r.Sign() == 0 {
			//log.Println("r.Sign() == 0")
			return false
		}

		e := hashToInt(hashClone, c)
		s = new(big.Int).Mul(priv.D, r)
		s.Add(s, e)
		s.Mul(s, inv)
		s.Mod(s, n)

		if s.Sign() == 0 {
			//log.Println("s.Sign() == 0")
			return false
		}

		return true
	}, nonce)

	//log.Println("enforce low S values, see bip62: 'low s values in signatures'");
	// enforce low S values, see bip62: 'low s values in signatures'
	nOverTwo := new(big.Int).Div(n, big.NewInt(2))
	if s.Cmp(nOverTwo) > 0 {
		s = new(big.Int).Sub(n, s)
	}

	return
}

// copied from crypto/ecdsa
func hashToInt(hash []byte, c elliptic.Curve) *big.Int {
	var hashClone = make([]byte, len(hash))
	copy(hashClone, hash)

	orderBits := c.Params().N.BitLen()
	orderBytes := (orderBits + 7) / 8
	if len(hashClone) > orderBytes {
		hashClone = hashClone[:orderBytes]
	}

	ret := new(big.Int).SetBytes(hashClone)
	excess := len(hashClone)*8 - orderBits
	if excess > 0 {
		ret.Rsh(ret, uint(excess))
	}
	return ret
}

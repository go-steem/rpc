/*
Package rfc6979 is an implementation of RFC 6979's deterministic DSA.
	Such signatures are compatible with standard Digital Signature Algorithm
	(DSA) and Elliptic Curve Digital Signature Algorithm (ECDSA) digital
	signatures and can be processed with unmodified verifiers, which need not be
	aware of the procedure described therein.  Deterministic signatures retain
	the cryptographic security features associated with digital signatures but
	can be more easily implemented in various environments, since they do not
	need access to a source of high-quality randomness.
(https://tools.ietf.org/html/rfc6979)
Provides functions similar to crypto/dsa and crypto/ecdsa.
*/
package rfc6979

import (
	"bytes"
	"crypto/hmac"
	"hash"
	"math/big"
	//"log"
	//"encoding/hex"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/binary"
	"math/rand"
)

//var one = big.NewInt(1)
var oneInitializer = []byte{0x01}

func RandStringBytes(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// https://tools.ietf.org/html/rfc6979#section-3.2
func generateSecret(priv *ecdsa.PrivateKey, alg func() hash.Hash, hash []byte, test func(*big.Int) bool, nonce int) error {
	var hashClone = make([]byte, len(hash))
	var err error
	copy(hashClone, hash)

	if nonce > 0 {
		nonceA := make([]byte, 4)
		binary.BigEndian.PutUint32(nonceA, uint32(nonce))
		hashClone = append(hashClone, nonceA...)
		hs := sha256.New()
		_, err = hs.Write(hashClone)
		if err != nil {
			return err
		}
		hashClone = hs.Sum(nil)
	}

	c := priv.PublicKey.Curve
	x := priv.D.Bytes()
	q := c.Params().N

	// Step B
	v := bytes.Repeat(oneInitializer, 32)

	// Step C (Go zeroes the all allocated memory)
	k := make([]byte, 32)

	// Step D

	m := append(append(append(v, 0x00), x...), hashClone...)

	k, err = HmacSHA256(m, k)
	if err != nil {
		return err
	}

	// Step E
	v, err = HmacSHA256(v, k)
	if err != nil {
		return err
	}

	// Step F
	k, err = HmacSHA256(append(append(append(v, 0x01), x...), hashClone...), k)
	if err != nil {
		return err
	}

	// Step G
	v, err = HmacSHA256(v, k)
	if err != nil {
		return err
	}

	// Step H1/H2a, ignored as tlen === qlen (256 bit)
	// Step H2b
	v, err = HmacSHA256(v, k)
	if err != nil {
		return err
	}

	var t = hashToInt(v, c)
	if err != nil {
		return err
	}

	// Step H3, repeat until T is within the interval [1, n - 1]
	for t.Sign() <= 0 || t.Cmp(q) >= 0 || !test(t) {

		k, err = HmacSHA256(append(v, 0x00), k)
		if err != nil {
			return err
		}

		v, err = HmacSHA256(v, k)
		if err != nil {
			return err
		}

		// Step H1/H2a, again, ignored as tlen === qlen (256 bit)
		// Step H2b again
		v, err = HmacSHA256(v, k)
		if err != nil {
			return err
		}

		t = hashToInt(v, c)
	}
	return nil
}

func HmacSHA256(m, k []byte) ([]byte, error) {
	mac := hmac.New(sha256.New, k)
	_, err := mac.Write(m)
	if err != nil {
		return []byte{}, err
	}
	expectedMAC := mac.Sum(nil)
	return expectedMAC, nil
}

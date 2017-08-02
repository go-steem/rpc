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
	"github.com/tendermint/go-crypto"
	"math/rand"
)

// mac returns an HMAC of the given key and message.
func mac(alg func() hash.Hash, k, m, buf []byte) []byte {
	h := hmac.New(alg, k)
	h.Write(m)
	return h.Sum(buf[:0])
}

// https://tools.ietf.org/html/rfc6979#section-2.3.2
func bits2int(in []byte, qlen int) *big.Int {
	vlen := len(in) * 8
	v := new(big.Int).SetBytes(in)
	if vlen > qlen {
		v = new(big.Int).Rsh(v, uint(vlen-qlen))
	}
	return v
}

// https://tools.ietf.org/html/rfc6979#section-2.3.3
func int2octets(v *big.Int, rolen int) []byte {
	out := v.Bytes()

	// pad with zeros if it's too short
	if len(out) < rolen {
		out2 := make([]byte, rolen)
		copy(out2[rolen-len(out):], out)
		return out2
	}

	// drop most significant bytes if it's too long
	if len(out) > rolen {
		out2 := make([]byte, rolen)
		copy(out2, out[len(out)-rolen:])
		return out2
	}

	return out
}

// https://tools.ietf.org/html/rfc6979#section-2.3.4
func bits2octets(in []byte, q *big.Int, qlen, rolen int) []byte {
	z1 := bits2int(in, qlen)
	z2 := new(big.Int).Sub(z1, q)
	if z2.Sign() < 0 {
		return int2octets(z1, rolen)
	}
	return int2octets(z2, rolen)
}

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
func generateSecret(priv *ecdsa.PrivateKey /*q, x *big.Int,*/, alg func() hash.Hash, hash []byte, test func(*big.Int) bool, nonce int) {
	//log.Println("priv=", priv.)
	var hash_clone = make([]byte, len(hash))
	copy(hash_clone, hash)

	//log.Println("before hash_clone=", hex.EncodeToString(hash_clone))
	if nonce > 0 {
		//nonce_str := RandStringBytes(nonce)
		nonce_a := make([]byte, 4)
		binary.BigEndian.PutUint32(nonce_a, uint32(nonce))
		hash_clone = append(hash_clone, nonce_a...)
		//log.Println("(before hash) hash_clone=", hex.EncodeToString(hash_clone), "nonce_str=", hex.EncodeToString(nonce_a))
		hash_clone = crypto.Sha256(hash_clone)
	}
	//log.Println("hash_clone=", hex.EncodeToString(hash_clone))

	c := priv.PublicKey.Curve
	//N := c.Params().N
	x := priv.D.Bytes()
	q := c.Params().N
	//x := privkey.Bytes()
	//alg := fastsha256.New

	//qlen := q.BitLen()
	//holen := alg().Size()

	//rolen := (qlen + 7) >> 3
	//bx := append(int2octets(x, rolen), bits2octets(hash, curve, rolen)...)

	//log.Println("bx=", hex.EncodeToString(bx))

	// Step B
	v := bytes.Repeat(oneInitializer, 32)

	// Step C (Go zeroes the all allocated memory)
	k := make([]byte, 32)

	// Step D
	//k = mac(alg, k, append(append(append(v, 0x00), bx...), hash... ))
	m := append(append(append(v, 0x00), x...), hash_clone...)
	//log.Println("m", hex.EncodeToString(m))
	//log.Println("k", hex.EncodeToString(k))
	k = HmacSHA256(m, k)
	//log.Println("Step D", hex.EncodeToString(k))

	// Step E
	//v = mac(alg, k, v)
	v = HmacSHA256(v, k)
	//log.Println("Step E", hex.EncodeToString(v))

	// Step F
	//k = mac(alg, k, append(append(append(v, 0x01), bx...), hash...))
	k = HmacSHA256(append(append(append(v, 0x01), x...), hash_clone...), k)
	//log.Println("Step F", hex.EncodeToString(k))

	// Step G
	//v = mac(alg, k, v)
	v = HmacSHA256(v, k)
	//log.Println("Step G", hex.EncodeToString(v))

	// Step H1/H2a, ignored as tlen === qlen (256 bit)
	// Step H2b
	v = HmacSHA256(v, k)
	//log.Println("Step H2b", hex.EncodeToString(v))

	//if (nonce.Cmp(big.NewInt(0)) != 0) {
	//	alg := sha256.New()
	//	alg.Write(hash)
	//	alg.Write(nonce.Bytes())
	//	hash = alg.Sum(nil)
	//}
	//
	//qlen := q.BitLen()
	//holen := alg().Size()
	//rolen := (qlen + 7) >> 3
	//bx := append(int2octets(x, rolen), bits2octets(hash, q, qlen, rolen)...)
	//
	//// Step B
	//v := bytes.Repeat([]byte{0x01}, holen)
	//
	//// Step C
	//k := bytes.Repeat([]byte{0x00}, holen)
	//
	//// Step D
	//k = mac(alg, k, append(append(v, 0x00), bx...), k)
	//
	//// Step E
	//v = mac(alg, k, v, v)
	//
	//// Step F
	//k = mac(alg, k, append(append(v, 0x01), bx...), k)
	//
	//// Step G
	//v = mac(alg, k, v, v)

	//////////////////////

	var T = hashToInt(v, c)
	//log.Println("T", hex.EncodeToString(T.Bytes()))

	// Step H3, repeat until T is within the interval [1, n - 1]
	for T.Sign() <= 0 || T.Cmp(q) >= 0 || !test(T) {

		//k = crypto.HmacSHA256(Buffer.concat([v, new Buffer([0])]), k);
		k = HmacSHA256(append(v, 0x00), k)

		//v = crypto.HmacSHA256(v, k);
		v = HmacSHA256(v, k)

		// Step H1/H2a, again, ignored as tlen === qlen (256 bit)
		// Step H2b again
		//v = crypto.HmacSHA256(v, k);
		v = HmacSHA256(v, k)

		//T = BigInteger.fromBuffer(v);
		T = hashToInt(v, c)

		//log.Println("T", hex.EncodeToString(T.Bytes()))
	}

	//return T;

	//// Step H
	//for {
	//	// Step H1
	//	var t []byte
	//
	//	// Step H2
	//	for len(t) < qlen/8 {
	//		v = mac(alg, k, v, v)
	//		t = append(t, v...)
	//	}
	//
	//	// Step H3
	//	secret := bits2int(t, qlen)
	//	log.Println("secret", hex.EncodeToString(secret.Bytes()))
	//	if secret.Cmp(one) >= 0 && secret.Cmp(q) < 0 && test(secret) {
	//		return
	//	}
	//	k = mac(alg, k, append(v, 0x00), k)
	//	v = mac(alg, k, v, v)
	//}
}

func HmacSHA256(m []byte, k []byte) []byte {
	//return mac(fastsha256.New, crypto.Sha256(m), crypto.Sha256(k))
	mac := hmac.New(sha256.New, k)
	mac.Write(m)
	expectedMAC := mac.Sum(nil)
	return expectedMAC
}

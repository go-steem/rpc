package client

import (
	crand "crypto/rand"
	"crypto/sha256"
	"math/rand"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

var src rand.Source

func init() {
	seed := time.Now().UnixNano()

	reader := crand.Reader
	i, err := crand.Prime(reader, 64)
	if err != nil {
		seed = seed ^ i.Int64()
	}

	src = rand.NewSource(seed)
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

//GenPassword allows to generate a 52-character password of the evil system.
func GenPassword() string {
	b := make([]byte, 51)
	for i, cache, remain := 51-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return "P" + string(b)
}

//GetPrivateKey generates a private key based on the specified parameters.
func GetPrivateKey(user, role, password string) string {
	hashSha256 := sha256.Sum256([]byte(user + role + password))
	pk := append([]byte{0x80}, hashSha256[:]...)
	chs := sha256.Sum256(pk)
	chs = sha256.Sum256(chs[:])
	b58 := append(pk, chs[:4]...)
	return base58.Encode(b58)
}

//GetPublicKey generates a public key based on the prefix and the private key.
func GetPublicKey(prefix, privatekey string) string {
	b58 := base58.Decode(privatekey)
	tpk := b58[:len(b58)-4]
	chs := b58[len(b58)-4:]
	nchs := sha256.Sum256(tpk)
	nchs = sha256.Sum256(nchs[:])
	if string(chs) != string(nchs[:4]) {
		return "Invalid WIF key (checksum miss-match)"
	}
	privKeyBytes := [32]byte{}
	copy(privKeyBytes[:], tpk[1:])
	priv, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes[:])
	chHash := ripemd160.New()
	chHash.Write(priv.PubKey().SerializeCompressed())
	nc := chHash.Sum(nil)
	pk := append(priv.PubKey().SerializeCompressed(), nc[:4]...)
	return prefix + base58.Encode(pk)
}

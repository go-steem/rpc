// +build !nosigning

package transactions

import (
	// Stdlib
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"time"

	// RPC
	"github.com/asuleymanov/golos-go/encoding/transaction"
	"github.com/asuleymanov/golos-go/types"

	// Vendor
	"github.com/pkg/errors"
)

type SignedTransaction struct {
	*types.Transaction
}

func NewSignedTransaction(tx *types.Transaction) *SignedTransaction {
	if tx.Expiration == nil {
		expiration := time.Now().Add(30 * time.Second).UTC()
		tx.Expiration = &types.Time{&expiration}
	}

	return &SignedTransaction{tx}
}

func (tx *SignedTransaction) Serialize() ([]byte, error) {
	var b bytes.Buffer
	encoder := transaction.NewEncoder(&b)

	if err := encoder.Encode(tx.Transaction); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (tx *SignedTransaction) Digest(chain *Chain) ([]byte, error) {
	var msgBuffer bytes.Buffer

	// Write the chain ID.
	rawChainID, err := hex.DecodeString(chain.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decode chain ID: %v", chain.ID)
	}

	if _, err := msgBuffer.Write(rawChainID); err != nil {
		return nil, errors.Wrap(err, "failed to write chain ID")
	}

	// Write the serialized transaction.
	rawTx, err := tx.Serialize()
	if err != nil {
		return nil, err
	}

	if _, err := msgBuffer.Write(rawTx); err != nil {
		return nil, errors.Wrap(err, "failed to write serialized transaction")
	}

	// Compute the digest.
	digest := sha256.Sum256(msgBuffer.Bytes())
	return digest[:], nil
}

// get rid of cgo and lsecp256k1
func (tx *SignedTransaction) Sign(privKeys [][]byte, chain *Chain) error {
	var buf bytes.Buffer
	chainid, _ := hex.DecodeString(chain.ID)
	//fmt.Println(tx.Operations[0])
	//fmt.Println(" ")
	tx_raw, _ := tx.Serialize()
	//fmt.Println(tx_raw)
	//fmt.Println(" ")
	buf.Write(chainid)
	buf.Write(tx_raw)
	data := buf.Bytes()
	//msg_sha := crypto.Sha256(buf.Bytes())

	var sigsHex []string

	for _, priv_b := range privKeys {
		sigBytes := tx.Sign_Single(priv_b, data)
		sigsHex = append(sigsHex, hex.EncodeToString(sigBytes))
	}

	tx.Transaction.Signatures = sigsHex
	return nil
}

//func (tx *SignedTransaction) Verify(pubKeys [][]byte, chain *Chain) (bool, error) {
//	// Compute the digest, again.
//	digest, err := tx.Digest(chain)
//	if err != nil {
//		return false, err
//	}
//
//	cDigest := C.CBytes(digest)
//	defer C.free(cDigest)
//
//	// Make sure to free memory.
//	cSigs := make([]unsafe.Pointer, 0, len(tx.Signatures))
//	defer func() {
//		for _, cSig := range cSigs {
//			C.free(cSig)
//		}
//	}()
//
//	// Collect verified public keys.
//	pubKeysFound := make([][]byte, len(pubKeys))
//	for i, signature := range tx.Signatures {
//		sig, err := hex.DecodeString(signature)
//		if err != nil {
//			return false, errors.Wrap(err, "failed to decode signature hex")
//		}
//
//		recoverParameter := sig[0] - 27 - 4
//		sig = sig[1:]
//
//		cSig := C.CBytes(sig)
//		cSigs = append(cSigs, cSig)
//
//		var publicKey [33]byte
//
//		code := C.verify_recoverable_signature(
//			(*C.uchar)(cDigest),
//			(*C.uchar)(cSig),
//			(C.int)(recoverParameter),
//			(*C.uchar)(&publicKey[0]),
//		)
//		if code == 1 {
//			pubKeysFound[i] = publicKey[:]
//		}
//	}
//
//	for i := range pubKeys {
//		if !bytes.Equal(pubKeysFound[i], pubKeys[i]) {
//			return false, nil
//		}
//	}
//	return true, nil
//}

// +build !nosigning

package transactions

import (
	// Stdlib
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"time"
	"unsafe"

	// RPC
	"github.com/asuleymanov/golos-go/encoding/transaction"
	"github.com/asuleymanov/golos-go/types"

	// Vendor
	"github.com/pkg/errors"
)

// #cgo LDFLAGS: -lsecp256k1
// #include <stdlib.h>
// #include "signing.h"
import "C"

type SignedTransaction struct {
	*types.Transaction
}

func NewSignedTransaction(tx *types.Transaction) *SignedTransaction {
	if tx.Expiration == nil {
		expiration := time.Now().Add(30 * time.Second)
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

func (tx *SignedTransaction) Sign(privKeys [][]byte, chain *Chain) error {
	digest, err := tx.Digest(chain)
	if err != nil {
		return err
	}

	// Sign.
	cDigest := C.CBytes(digest)
	defer C.free(cDigest)

	cKeys := make([]unsafe.Pointer, 0, len(privKeys))
	for _, key := range privKeys {
		cKeys = append(cKeys, C.CBytes(key))
	}
	defer func() {
		for _, cKey := range cKeys {
			C.free(cKey)
		}
	}()

	sigs := make([][]byte, 0, len(privKeys))
	for _, cKey := range cKeys {
		var (
			signature [64]byte
			recid     C.int
		)

		code := C.sign_transaction(
			(*C.uchar)(cDigest), (*C.uchar)(cKey), (*C.uchar)(&signature[0]), &recid)
		if code == 0 {
			return errors.New("sign_transaction returned 0")
		}

		sig := make([]byte, 65)
		sig[0] = byte(recid)
		copy(sig[1:], signature[:])

		sigs = append(sigs, sig)
	}

	// Set the signature array in the transaction.
	sigsHex := make([]string, 0, len(sigs))
	for _, sig := range sigs {
		sigsHex = append(sigsHex, hex.EncodeToString(sig))
	}

	tx.Transaction.Signatures = sigsHex
	return nil
}

func (tx *SignedTransaction) Verify(pubKeys [][]byte, chain *Chain) (bool, error) {
	// Compute the digest, again.
	digest, err := tx.Digest(chain)
	if err != nil {
		return false, err
	}

	cDigest := C.CBytes(digest)
	defer C.free(cDigest)

	// Make sure to free memory.
	cSigs := make([]unsafe.Pointer, 0, len(tx.Signatures))
	defer func() {
		for _, cSig := range cSigs {
			C.free(cSig)
		}
	}()

	// Collect verified public keys.
	pubKeysFound := make([][]byte, len(pubKeys))
	for i, signature := range tx.Signatures {
		sig, err := hex.DecodeString(signature)
		if err != nil {
			return false, errors.Wrap(err, "failed to decode signature hex")
		}

		recoverParameter := sig[0] - 27 - 4
		sig = sig[1:]

		cSig := C.CBytes(sig)
		cSigs = append(cSigs, cSig)

		var publicKey [33]byte

		code := C.verify_recoverable_signature(
			(*C.uchar)(cDigest),
			(*C.uchar)(cSig),
			(C.int)(recoverParameter),
			(*C.uchar)(&publicKey[0]),
		)
		if code == 1 {
			pubKeysFound[i] = publicKey[:]
		}
	}

	for i := range pubKeys {
		if !bytes.Equal(pubKeysFound[i], pubKeys[i]) {
			return false, nil
		}
	}
	return true, nil
}

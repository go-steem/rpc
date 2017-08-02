package types

import (
	// Stdlib
	"bytes"
	"encoding/hex"
	"testing"
	"time"

	// RPC
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

func TestTransaction_MarshalTransaction(t *testing.T) {
	// The result we expect.
	expected := "bd8c5fe26f45f179a8570100057865726f63057865726f6306706973746f6e102700"

	// Prepare the transaction.
	expiration := time.Date(2016, 8, 8, 12, 24, 17, 0, time.UTC)
	tx := Transaction{
		RefBlockNum:    36029,
		RefBlockPrefix: 1164960351,
		Expiration:     &Time{&expiration},
	}
	tx.PushOperation(&VoteOperation{
		Voter:    "xeroc",
		Author:   "xeroc",
		Permlink: "piston",
		Weight:   10000,
	})

	// Marshal the transaction.
	var b bytes.Buffer
	encoder := transaction.NewEncoder(&b)

	if err := tx.MarshalTransaction(encoder); err != nil {
		t.Error(err)
	}
	got := hex.EncodeToString(b.Bytes())

	// Compare that we got with what we expect to get.
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

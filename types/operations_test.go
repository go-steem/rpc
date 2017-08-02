package types

import (
	// Stdlib
	"bytes"
	"encoding/hex"
	"testing"

	// RPC
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

func TestVoteOperation_MarshalTransaction(t *testing.T) {
	op := &VoteOperation{
		Voter:    "xeroc",
		Author:   "xeroc",
		Permlink: "piston",
		Weight:   10000,
	}

	expectedHex := "00057865726f63057865726f6306706973746f6e1027"

	var b bytes.Buffer
	encoder := transaction.NewEncoder(&b)

	if err := encoder.Encode(op); err != nil {
		t.Error(err)
	}

	serializedHex := hex.EncodeToString(b.Bytes())

	if serializedHex != expectedHex {
		t.Errorf("expected %v, got %v", expectedHex, serializedHex)
	}
}

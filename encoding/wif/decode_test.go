package wif

import (
	// Stdlib
	"encoding/hex"
	"testing"
)

func TestDecode(t *testing.T) {
	for _, d := range data {
		privKey, err := Decode(d.WIF)
		if err != nil {
			t.Error(err)
		}

		expected := d.PrivateKeyHex
		got := hex.EncodeToString(privKey)

		if got != expected {
			t.Errorf("expected %v, got %v", expected, got)
		}
	}
}

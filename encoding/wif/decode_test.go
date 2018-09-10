package wif

import (
	// Stdlib
	"encoding/hex"
	"testing"
)

type testData struct {
	WIF           string
	PrivateKeyHex string
}

var data = []testData{
	{
		WIF:           "5JWHY5DxTF6qN5grTtChDCYBmWHfY9zaSsw4CxEKN5eZpH9iBma",
		PrivateKeyHex: "5ad2b8df2c255d4a2996ee7d065e013e1bbb35c075ee6e5208aca44adc9a9d4c",
	},
	{
		WIF:           "5KPipdRzoxrp6dDqsBfMD6oFZG356trVHV5QBGx3rABs1zzWWs8",
		PrivateKeyHex: "cf9d6121ed458f24ea456ad7ff700da39e86688988cfe5c6ed6558642cf1e32f",
	},
}

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

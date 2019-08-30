package types

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//Asset type from parameter JSON
type Asset struct {
	Amount float64
	Symbol string
}

//UnmarshalJSON unpacking the JSON parameter in the Asset type.
func (op *Asset) UnmarshalJSON(data []byte) error {
	str, errUnq := strconv.Unquote(string(data))
	if errUnq != nil {
		return errUnq
	}
	param := strings.Split(str, " ")

	s, errpf := strconv.ParseFloat(param[0], 64)
	if errpf != nil {
		return errpf
	}

	op.Amount = s
	op.Symbol = param[1]

	return nil
}

//MarshalJSON function for packing the Asset type in JSON.
func (op *Asset) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.String())
}

//MarshalTransaction is a function of converting type Asset to bytes.
func (op *Asset) MarshalTransaction(encoder *transaction.Encoder) error {
	ans, err := json.Marshal(op)
	if err != nil {
		return err
	}

	str, err := strconv.Unquote(string(ans))
	if err != nil {
		return err
	}
	return encoder.EncodeMoney(str)
}

//String function convert type Asset to string.
func (op *Asset) String() string {
	var ammf string
	if op.Symbol != "SHARES" {
		ammf = strconv.FormatFloat(op.Amount, 'f', 3, 64)
	} else {
		ammf = strconv.FormatFloat(op.Amount, 'f', 6, 64)
	}
	return ammf + " " + op.Symbol
}

//StringAmount function convert type Asset.Amount to string.
func (op *Asset) StringAmount() string {
	return strconv.FormatFloat(op.Amount, 'f', 3, 64)
}

package types

import (
	"encoding/json"
	"strconv"

	"github.com/asuleymanov/steem-go/encoding/transaction"
	"github.com/pkg/errors"
)

//ContentMetadata type from parameter JSON
type ContentMetadata map[string]interface{}

//UnmarshalJSON unpacking the JSON parameter in the ContentMetadata type.
func (op *ContentMetadata) UnmarshalJSON(p []byte) error {
	var raw map[string]interface{}

	str, errUnq := strconv.Unquote(string(p))
	if errUnq != nil {
		return errUnq
	}
	if str == "" {
		return nil
	}

	if err := json.Unmarshal([]byte(str), &raw); err != nil {
		return errors.Wrap(err, "ERROR: ContentMedata unmarshal error")
	}

	*op = raw

	return nil
}

//MarshalJSON function for packing the ContentMetadata type in JSON.
func (op *ContentMetadata) MarshalJSON() ([]byte, error) {
	ans, err := json.Marshal(*op)
	if err != nil {
		return []byte{}, err
	}
	return []byte(strconv.Quote(string(ans))), nil
}

//MarshalTransaction is a function of converting type ContentMetadata to bytes.
func (op *ContentMetadata) MarshalTransaction(encoder *transaction.Encoder) error {
	ans, err := json.Marshal(op)
	if err != nil {
		return err
	}

	str, err := strconv.Unquote(string(ans))
	if err != nil {
		return err
	}

	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeString(str)
	return enc.Err()
}

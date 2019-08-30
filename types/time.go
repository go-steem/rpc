package types

import (
	"time"

	"github.com/asuleymanov/steem-go/encoding/transaction"
)

const layout = `"2006-01-02T15:04:05"`

//Time type from parameter JSON
type Time struct {
	*time.Time
}

//MarshalJSON function for packing the Time type in JSON.
func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(layout)), nil
}

//UnmarshalJSON unpacking the JSON parameter in the Time type.
func (t *Time) UnmarshalJSON(data []byte) error {
	parsed, err := time.ParseInLocation(layout, string(data), time.UTC)
	if err != nil {
		return err
	}
	t.Time = &parsed
	return nil
}

//MarshalTransaction is a function of converting type Time to bytes.
func (t *Time) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.Encode(uint32(t.Time.Unix()))
}

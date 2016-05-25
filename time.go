package rpc

import (
	"time"
)

type Time struct {
	*time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	parsed, err := time.Parse(`"2006-01-02T15:04:05"`, string(data))
	if err != nil {
		return err
	}
	t.Time = &parsed
	return nil
}

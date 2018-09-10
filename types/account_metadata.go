package types

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/asuleymanov/steem-go/encoding/transaction"
	"github.com/pkg/errors"
)

//AccountMetadata type from parameter JSON
type AccountMetadata struct {
	Profile ProfileJSON `json:"profile,omitempty"`
}

type rawAccountMetadata struct {
	Profile ProfileJSON `json:"profile,omitempty"`
}

//ProfileJSON additional structure for the AccountMetadata type.
type ProfileJSON struct {
	Name         string `json:"name,omitempty"`
	ProfileImage string `json:"profile_image,omitempty"`
	CoverImage   string `json:"cover_image,omitempty"`
	Gender       string `json:"gender,omitempty"`
	About        string `json:"about,omitempty"`
	Location     string `json:"location,omitempty"`
	Website      string `json:"website,omitempty"`
}

//UnmarshalJSON unpacking the JSON parameter in the AccountMetadata type.
func (op *AccountMetadata) UnmarshalJSON(p []byte) error {
	var raw rawAccountMetadata

	str, errUnq := strconv.Unquote(string(p))
	if errUnq != nil {
		return errUnq
	}
	if str == "" {
		return nil
	}

	if err := json.Unmarshal([]byte(str), &raw); err != nil {
		return errors.Wrap(err, "ERROR: AccountMedata unmarshal error")
	}

	op.Profile = ProfileJSON{
		Name:         raw.Profile.Name,
		ProfileImage: raw.Profile.ProfileImage,
		CoverImage:   raw.Profile.CoverImage,
		Gender:       raw.Profile.Gender,
		About:        raw.Profile.About,
		Location:     raw.Profile.Location,
		Website:      raw.Profile.Website,
	}
	return nil
}

//MarshalJSON function for packing the AccountMetadata type in JSON.
func (op *AccountMetadata) MarshalJSON() ([]byte, error) {
	ans, err := json.Marshal(&rawAccountMetadata{
		Profile: ProfileJSON{
			Name:         op.Profile.Name,
			ProfileImage: op.Profile.ProfileImage,
			CoverImage:   op.Profile.CoverImage,
			Gender:       op.Profile.Gender,
			About:        op.Profile.About,
			Location:     op.Profile.Location,
			Website:      op.Profile.Website,
		},
	})
	if err != nil {
		return []byte{}, err
	}
	return []byte(strconv.Quote(string(ans))), nil
}

//MarshalTransaction is a function of converting type AccountMetadata to bytes.
func (op *AccountMetadata) MarshalTransaction(encoder *transaction.Encoder) error {
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

//String function convert type AccountMetadata to string.
func (op *AccountMetadata) String() string {
	return fmt.Sprintf("%#v", op)
}

package client

import (
	// Stdlib
	"log"

	// Vendor
	"github.com/pkg/errors"
)

//We check whether there is a voter on the list of those who have already voted
func (api *Client) Verify_Voter(author, permlink, voter string) bool {
	ans, err := api.Rpc.Database.GetActiveVotes(author, permlink)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Voter: "))
		return false
	} else {
		for _, v := range ans {
			if v.Voter == voter {
				return true
			}
		}
		return false
	}
}

//We check whether there are voted
func (api *Client) Verify_Votes(author, permlink string) bool {
	ans, err := api.Rpc.Database.GetActiveVotes(author, permlink)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Votes: "))
		return false
	} else {
		if len(ans) > 0 {
			return true
		} else {
			return false
		}
	}
}

func (api *Client) Verify_Comments(author, permlink string) bool {
	ans, err := api.Rpc.Database.GetContentReplies(author, permlink)
	if err != nil {
		log.Println(errors.Wrapf(err, "Error Verify Comments: "))
		return false
	} else {
		if len(ans) > 0 {
			return true
		} else {
			return false
		}
	}
}

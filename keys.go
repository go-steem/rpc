package client

import (
	"github.com/asuleymanov/steem-go/encoding/wif"
	"github.com/asuleymanov/steem-go/types"
	"github.com/pkg/errors"
)

var (
	//OpTypeKey include a description of the operation and the key needed to sign it
	OpTypeKey = make(map[types.OpType][]string)
)

//Keys is used as a keystroke for a specific user.
//Only a few keys can be set.
type Keys struct {
	PKey []string
	AKey []string
	OKey []string
	MKey []string
}

func init() {
	OpTypeKey["vote"] = []string{"posting"}
	OpTypeKey["content"] = []string{"posting"}
	OpTypeKey["transfer"] = []string{"active"}
	OpTypeKey["transfer_to_vesting"] = []string{"active"}
	OpTypeKey["withdraw_vesting"] = []string{"active"}
	OpTypeKey["account_update"] = []string{"active"}
	OpTypeKey["witness_update"] = []string{"active"}
	OpTypeKey["account_witness_vote"] = []string{"posting"}
	OpTypeKey["account_witness_proxy"] = []string{"posting"}
	OpTypeKey["delete_content"] = []string{"posting"}
	OpTypeKey["custom"] = []string{"posting"}
	OpTypeKey["set_withdraw_vesting_route"] = []string{"active"}
	OpTypeKey["request_account_recovery"] = []string{"active"}
	OpTypeKey["recover_account"] = []string{"owner"}
	OpTypeKey["change_recovery_account"] = []string{"owner"}
	OpTypeKey["escrow_transfer"] = []string{"active"}
	OpTypeKey["escrow_dispute"] = []string{"active"}
	OpTypeKey["escrow_release"] = []string{"active"}
	OpTypeKey["escrow_approve"] = []string{"active"}
	OpTypeKey["delegate_vesting_shares"] = []string{"active"}
	OpTypeKey["account_create"] = []string{"active"}
	OpTypeKey["account_metadata"] = []string{"posting"}
	OpTypeKey["proposal_create"] = []string{"active"}
	OpTypeKey["proposal_update"] = []string{"active"}
	OpTypeKey["proposal_delete"] = []string{"active"}
	OpTypeKey["chain_properties_update"] = []string{"active"}
	OpTypeKey["committee_worker_create_request"] = []string{"posting"}
	OpTypeKey["committee_worker_cancel_request"] = []string{"posting"}
	OpTypeKey["committee_vote_request"] = []string{"posting"}
	OpTypeKey["create_invite"] = []string{"active"}
	OpTypeKey["claim_invite_balance"] = []string{"active"}
	OpTypeKey["invite_registration"] = []string{"active"}
	OpTypeKey["versioned_chain_properties_update"] = []string{"active"}
	OpTypeKey["award"] = []string{"posting"}
	OpTypeKey["set_paid_subscription"] = []string{"active"}
	OpTypeKey["paid_subscribe"] = []string{"active"}
}

//SetKeys you can specify keys for signing transactions.
func (client *Client) SetKeys(keys *Keys) {
	client.CurrentKeys = keys
}

//SigningKeys returns the key from the CurrentKeys
func (client *Client) SigningKeys(trx types.Operation) ([][]byte, error) {
	var keys [][]byte

	if client.CurrentKeys == nil {
		return nil, errors.New("Client Keys not initialized. Use SetKeys method")
	}

	opKeys := OpTypeKey[trx.Type()]
	for _, val := range opKeys {
		switch val {
		case "posting":
			for _, keyStr := range client.CurrentKeys.PKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Posting Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case "active":
			for _, keyStr := range client.CurrentKeys.AKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Active Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case "owner":
			for _, keyStr := range client.CurrentKeys.OKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Owner Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case "memo":
			for _, keyStr := range client.CurrentKeys.MKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Memo Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		}
	}

	return keys, nil
}

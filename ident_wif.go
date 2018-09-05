package client

import (
	"github.com/pkg/errors"
	"github.com/asuleymanov/steem-go/encoding/wif"
	"github.com/asuleymanov/steem-go/types"
)

var (
	OpTypeKey = make(map[types.OpType][]string)
)

func init() {
	OpTypeKey["vote"] = []string{"posting"}
	OpTypeKey["comment"] = []string{"posting"}
	OpTypeKey["transfer"] = []string{"active"}
	OpTypeKey["transfer_to_vesting"] = []string{"active"}
	OpTypeKey["withdraw_vesting"] = []string{"active"}
	OpTypeKey["limit_order_create"] = []string{"active"}
	OpTypeKey["limit_order_cancel"] = []string{"active"}
	OpTypeKey["feed_publish"] = []string{"active"}
	OpTypeKey["convert"] = []string{"active"}
	OpTypeKey["account_create"] = []string{"active"}
	OpTypeKey["account_update"] = []string{"active"}
	OpTypeKey["witness_update"] = []string{"active"}
	OpTypeKey["account_witness_vote"] = []string{"active"}
	OpTypeKey["account_witness_proxy"] = []string{"active"}
	OpTypeKey["pow"] = []string{"active"}
	OpTypeKey["custom"] = []string{"active"}
	OpTypeKey["report_over_production"] = []string{"posting"}
	OpTypeKey["delete_comment"] = []string{"posting"}
	OpTypeKey["custom_json"] = []string{"posting"}
	OpTypeKey["comment_options"] = []string{"posting"}
	OpTypeKey["set_withdraw_vesting_route"] = []string{"active"}
	OpTypeKey["limit_order_create2"] = []string{"active"}
	OpTypeKey["challenge_authority"] = []string{"posting"}
	OpTypeKey["prove_authority"] = []string{"active"}
	OpTypeKey["request_account_recovery"] = []string{"active"}
	OpTypeKey["recover_account"] = []string{"owner"}
	OpTypeKey["change_recovery_account"] = []string{"owner"}
	OpTypeKey["escrow_transfer"] = []string{"active"}
	OpTypeKey["escrow_dispute"] = []string{"active"}
	OpTypeKey["escrow_release"] = []string{"active"}
	OpTypeKey["pow2"] = []string{"active"}
	OpTypeKey["escrow_approve"] = []string{"active"}
	OpTypeKey["transfer_to_savings"] = []string{"active"}
	OpTypeKey["transfer_from_savings"] = []string{"active"}
	OpTypeKey["cancel_transfer_from_savings"] = []string{"active"}
	OpTypeKey["custom_binary"] = []string{"posting"}
	OpTypeKey["decline_voting_rights"] = []string{"owner"}
	OpTypeKey["reset_account"] = []string{"active"}
	OpTypeKey["set_reset_account"] = []string{"posting"}
	OpTypeKey["claim_reward_balance"] = []string{"posting"}
	OpTypeKey["delegate_vesting_shares"] = []string{"active"}
	OpTypeKey["account_create_with_delegation"] = []string{"active"}
	OpTypeKey["fill_convert_request"] = []string{"active"}
	OpTypeKey["author_reward"] = []string{"posting"}
	OpTypeKey["curation_reward"] = []string{"posting"}
	OpTypeKey["comment_reward"] = []string{"posting"}
	OpTypeKey["liquidity_reward"] = []string{"active"}
	OpTypeKey["interest"] = []string{"active"}
	OpTypeKey["fill_vesting_withdraw"] = []string{"active"}
	OpTypeKey["fill_order"] = []string{"posting"}
	OpTypeKey["shutdown_witness"] = []string{"posting"}
	OpTypeKey["fill_transfer_from_savings"] = []string{"posting"}
	OpTypeKey["hardfork"] = []string{"posting"}
	OpTypeKey["comment_payout_update"] = []string{"posting"}
	OpTypeKey["return_vesting_delegation"] = []string{"posting"}
	OpTypeKey["comment_benefactor_reward"] = []string{"posting"}

}

//SigningKeys returns the key from the CurrentKeys
func (client *Client) SigningKeys(trx types.Operation) ([][]byte, error) {
	var keys [][]byte

	if client.CurrentKeys == nil {
		return nil, errors.New("client Keys not initialized. Use SetKeys method")
	}

	opKeys := OpTypeKey[trx.Type()]
	for _, val := range opKeys {
		switch {
		case val == "posting":
			for _, keyStr := range client.CurrentKeys.PKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Posting Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case val == "active":
			for _, keyStr := range client.CurrentKeys.AKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Active Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case val == "owner":
			for _, keyStr := range client.CurrentKeys.OKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Owner Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case val == "memo":
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

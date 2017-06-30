package client

import (
	"log"

	"github.com/pkg/errors"

	"github.com/asuleymanov/golos-go/encoding/wif"
	"github.com/asuleymanov/golos-go/types"
)

var (
	OpTypeKey = make(map[types.OpType][]string)
)

func init() {
	OpTypeKey["vote"] = []string{"posting"}
	OpTypeKey["comment"] = []string{"posting"}
	OpTypeKey["transfer"] = []string{"posting"}
	OpTypeKey["transfer_to_vesting"] = []string{"posting"}
	OpTypeKey["withdraw_vesting"] = []string{"posting"}
	OpTypeKey["limit_order_create"] = []string{"posting"}
	OpTypeKey["limit_order_cancel"] = []string{"posting"}
	OpTypeKey["feed_publish"] = []string{"posting"}
	OpTypeKey["convert"] = []string{"posting"}
	OpTypeKey["account_create"] = []string{"posting"}
	OpTypeKey["account_update"] = []string{"posting"}
	OpTypeKey["witness_update"] = []string{"posting"}
	OpTypeKey["account_witness_vote"] = []string{"posting"}
	OpTypeKey["account_witness_proxy"] = []string{"posting"}
	OpTypeKey["pow"] = []string{"posting"}
	OpTypeKey["custom"] = []string{"posting"}
	OpTypeKey["report_over_production"] = []string{"posting"}
	OpTypeKey["delete_comment"] = []string{"posting"}
	OpTypeKey["custom_json"] = []string{"posting"}
	OpTypeKey["comment_options"] = []string{"posting"}
	OpTypeKey["set_withdraw_vesting_route"] = []string{"posting"}
	OpTypeKey["limit_order_create2"] = []string{"posting"}
	OpTypeKey["challenge_authority"] = []string{"posting"}
	OpTypeKey["prove_authority"] = []string{"posting"}
	OpTypeKey["request_account_recovery"] = []string{"posting"}
	OpTypeKey["recover_account"] = []string{"posting"}
	OpTypeKey["change_recovery_account"] = []string{"posting"}
	OpTypeKey["escrow_transfer"] = []string{"posting"}
	OpTypeKey["escrow_dispute"] = []string{"posting"}
	OpTypeKey["escrow_release"] = []string{"posting"}
	OpTypeKey["pow2"] = []string{"posting"}
	OpTypeKey["escrow_approve"] = []string{"posting"}
	OpTypeKey["transfer_to_savings"] = []string{"posting"}
	OpTypeKey["transfer_from_savings"] = []string{"posting"}
	OpTypeKey["cancel_transfer_from_savings"] = []string{"posting"}
	OpTypeKey["custom_binary"] = []string{"posting"}
	OpTypeKey["decline_voting_rights"] = []string{"posting"}
	OpTypeKey["reset_account"] = []string{"posting"}
	OpTypeKey["set_reset_account"] = []string{"posting"}
	OpTypeKey["claim_reward_balance"] = []string{"posting"}
	OpTypeKey["delegate_vesting_shares"] = []string{"posting"}
	OpTypeKey["account_create_with_delegation"] = []string{"posting"}
	OpTypeKey["fill_convert_request"] = []string{"posting"}
	OpTypeKey["author_reward"] = []string{"posting"}
	OpTypeKey["curation_reward"] = []string{"posting"}
	OpTypeKey["comment_reward"] = []string{"posting"}
	OpTypeKey["liquidity_reward"] = []string{"posting"}
	OpTypeKey["interest"] = []string{"posting"}
	OpTypeKey["fill_vesting_withdraw"] = []string{"posting"}
	OpTypeKey["fill_order"] = []string{"posting"}
	OpTypeKey["shutdown_witness"] = []string{"posting"}
	OpTypeKey["fill_transfer_from_savings"] = []string{"posting"}
	OpTypeKey["hardfork"] = []string{"posting"}
	OpTypeKey["comment_payout_update"] = []string{"posting"}
	OpTypeKey["return_vesting_delegation"] = []string{"posting"}
	OpTypeKey["comment_benefactor_reward"] = []string{"posting"}

}

func (api *Golos) Signing_Keys(trx types.Operation) [][]byte {
	var keys [][]byte
	op_keys := OpTypeKey[trx.Type()]
	for _, val := range op_keys {
		switch {
		case val == "posting":
			privKey, err := wif.Decode(string([]byte(api.User.PKey)))
			if err != nil {
				log.Println(errors.Wrapf(err, "Error decode Key: "))
			}
			keys = append(keys, privKey)
		case val == "active":
			privKey, err := wif.Decode(string([]byte(api.User.AKey)))
			if err != nil {
				log.Println(errors.Wrapf(err, "Error decode Key: "))
			}
			keys = append(keys, privKey)
		case val == "owner":
			privKey, err := wif.Decode(string([]byte(api.User.OKey)))
			if err != nil {
				log.Println(errors.Wrapf(err, "Error decode Key: "))
			}
			keys = append(keys, privKey)
		case val == "memo":
			privKey, err := wif.Decode(string([]byte(api.User.MKey)))
			if err != nil {
				log.Println(errors.Wrapf(err, "Error decode Key: "))
			}
			keys = append(keys, privKey)
		}
	}
	return keys
}

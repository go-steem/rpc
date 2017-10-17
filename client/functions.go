package client

import (
	// Stdlib
	"log"
	"strconv"
	"time"

	// Vendor
	"github.com/pkg/errors"

	// RPC
	"github.com/asuleymanov/golos-go/encoding/wif"
	"github.com/asuleymanov/golos-go/transactions"
	"github.com/asuleymanov/golos-go/translit"
	"github.com/asuleymanov/golos-go/types"
)

func (api *Client) Vote(user_name, author_name, permlink string, weight int) error {
	if weight > 10000 {
		weight = 10000
	}
	if api.Verify_Voter(author_name, permlink, user_name) {
		return errors.New("The voter is on the list")
	}
	tx := &types.VoteOperation{
		Voter:    user_name,
		Author:   author_name,
		Permlink: permlink,
		Weight:   types.Int16(weight),
	}
	resp, err := api.Send_Trx(user_name, tx)
	if err != nil {
		return errors.Wrapf(err, "Error Vote: ")
	} else {
		log.Println("[Vote] Block -> ", resp.BlockNum, " User -> ", user_name)
		return nil
	}
}

func (api *Client) Comment(user_name, author_name, ppermlink, body string) error {
	times, _ := strconv.Unquote(time.Now().Add(30 * time.Second).UTC().Format(fdt))
	permlink := "re-" + author_name + "-" + ppermlink + "-" + times
	tx := &types.CommentOperation{
		ParentAuthor:   author_name,
		ParentPermlink: ppermlink,
		Author:         user_name,
		Permlink:       permlink,
		Title:          "",
		Body:           body,
		JsonMetadata:   "{\"app\":\"golos-go(go-steem)\"}",
	}
	resp, err := api.Send_Trx(user_name, tx)
	if err != nil {
		return errors.Wrapf(err, "Error Comment: ")
	} else {
		log.Println("[Comment] Block -> ", resp.BlockNum, " User -> ", user_name)
		return nil
	}
}

func (api *Client) Comment_Vote(user_name, author_name, ppermlink, body string, weight_post int) error {
	if weight_post > 10000 {
		weight_post = 10000
	}
	times, _ := strconv.Unquote(time.Now().Add(30 * time.Second).UTC().Format(fdt))
	permlink := "re-" + author_name + "-" + ppermlink + "-" + times
	var trx []types.Operation
	txc := &types.CommentOperation{
		ParentAuthor:   author_name,
		ParentPermlink: ppermlink,
		Author:         user_name,
		Permlink:       permlink,
		Title:          "",
		Body:           body,
		JsonMetadata:   "{\"app\":\"golos-go(go-steem)\"}",
	}
	trx = append(trx, txc)

	if !api.Verify_Voter(author_name, permlink, user_name) {
		txv := &types.VoteOperation{
			Voter:    user_name,
			Author:   author_name,
			Permlink: ppermlink,
			Weight:   types.Int16(weight_post),
		}
		trx = append(trx, txv)
	}

	resp, err := api.Send_Arr_Trx(user_name, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Comment and Vote: ")
	} else {
		log.Println("[Comment and Vote] Block -> ", resp.BlockNum, " User -> ", user_name)
		return nil
	}
}

func (api *Client) DeleteComment(author_name, permlink string) error {
	if api.Verify_Votes(author_name, permlink) {
		return errors.New("You can not delete already there are voted")
	}
	if api.Verify_Comments(author_name, permlink) {
		return errors.New("You can not delete already have comments")
	}
	tx := &types.DeleteCommentOperation{
		Author:   author_name,
		Permlink: permlink,
	}
	resp, err := api.Send_Trx(author_name, tx)
	if err != nil {
		return errors.Wrapf(err, "Error Delete Comment: ")
	} else {
		log.Println("[Delete Comment] Block -> ", resp.BlockNum, " User -> ", author_name)
		return nil
	}
}

func (api *Client) Post(author_name, title, body, permlink, ptag, post_image string, tags []string) error {
	if permlink == "" {
		permlink = translit.EncodeTitle(title)
	} else {
		permlink = translit.EncodeTitle(permlink)
	}
	tag := translit.EncodeTags(tags)
	if ptag == "" {
		ptag = translit.EncodeTag(tags[0])
	} else {
		ptag = translit.EncodeTag(ptag)
	}

	json_meta := "{\"tags\":["
	for k, v := range tag {
		if k != len(tags)-1 {
			json_meta = json_meta + "\"" + v + "\","
		} else {
			json_meta = json_meta + "\"" + v + "\"]"
		}
	}
	if post_image != "" {
		json_meta = json_meta + ",\"image\":\"" + post_image + "\""
	}
	json_meta = json_meta + ",\"app\":\"golos-go(go-steem)\"}"

	tx := &types.CommentOperation{
		ParentAuthor:   "",
		ParentPermlink: ptag,
		Author:         author_name,
		Permlink:       permlink,
		Title:          title,
		Body:           body,
		JsonMetadata:   json_meta,
	}

	resp, err := api.Send_Trx(author_name, tx)
	if err != nil {
		return errors.Wrapf(err, "Error Post: ")
	} else {
		log.Println("[Post] Block -> ", resp.BlockNum, " User -> ", author_name)
		return nil
	}
}

func (api *Client) Post_Vote(author_name, title, body, permlink, ptag, post_image string, tags []string, weight_post int) error {
	if weight_post > 10000 {
		weight_post = 10000
	}
	if permlink == "" {
		permlink = translit.EncodeTitle(title)
	} else {
		permlink = translit.EncodeTitle(permlink)
	}
	tag := translit.EncodeTags(tags)
	if ptag == "" {
		ptag = translit.EncodeTag(tags[0])
	} else {
		ptag = translit.EncodeTag(ptag)
	}

	json_meta := "{\"tags\":["
	for k, v := range tag {
		if k != len(tags)-1 {
			json_meta = json_meta + "\"" + v + "\","
		} else {
			json_meta = json_meta + "\"" + v + "\"]"
		}
	}
	if post_image != "" {
		json_meta = json_meta + ",\"image\":\"" + post_image + "\""
	}
	json_meta = json_meta + ",\"app\":\"golos-go(go-steem)\"}"

	var trx []types.Operation
	txp := &types.CommentOperation{
		ParentAuthor:   "",
		ParentPermlink: ptag,
		Author:         author_name,
		Permlink:       permlink,
		Title:          title,
		Body:           body,
		JsonMetadata:   json_meta,
	}
	trx = append(trx, txp)

	txv := &types.VoteOperation{
		Voter:    author_name,
		Author:   author_name,
		Permlink: permlink,
		Weight:   types.Int16(weight_post),
	}
	trx = append(trx, txv)

	resp, err := api.Send_Arr_Trx(author_name, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Post and Vote: ")
	} else {
		log.Println("[Post and Vote] Block -> ", resp.BlockNum, " User -> ", author_name)
		return nil
	}
}

func (api *Client) Post_Options(author_name, title, body, permlink, ptag, post_image string, tags []string, percent uint16, votes, curation bool) error {
	if permlink == "" {
		permlink = translit.EncodeTitle(title)
	} else {
		permlink = translit.EncodeTitle(permlink)
	}
	tag := translit.EncodeTags(tags)
	if ptag == "" {
		ptag = translit.EncodeTag(tags[0])
	} else {
		ptag = translit.EncodeTag(ptag)
	}
	MAP := "1000000.000 GBG"
	PSD := percent
	if percent == 0 {
		MAP = "0.000 GBG"
	} else if percent == 50 {
		PSD = 10000
	} else {
		PSD = 0
	}

	json_meta := "{\"tags\":["
	for k, v := range tag {
		if k != len(tags)-1 {
			json_meta = json_meta + "\"" + v + "\","
		} else {
			json_meta = json_meta + "\"" + v + "\"]"
		}
	}
	if post_image != "" {
		json_meta = json_meta + ",\"image\":\"" + post_image + "\""
	}
	json_meta = json_meta + ",\"app\":\"golos-go(go-steem)\"}"

	var trx []types.Operation
	txp := &types.CommentOperation{
		ParentAuthor:   "",
		ParentPermlink: ptag,
		Author:         author_name,
		Permlink:       permlink,
		Title:          title,
		Body:           body,
		JsonMetadata:   json_meta,
	}
	trx = append(trx, txp)

	txo := &types.CommentOptionsOperation{
		Author:               author_name,
		Permlink:             permlink,
		MaxAcceptedPayout:    MAP,
		PercentSteemDollars:  PSD,
		AllowVotes:           votes,
		AllowCurationRewards: curation,
		Extensions:           []interface{}{},
	}
	trx = append(trx, txo)

	resp, err := api.Send_Arr_Trx(author_name, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Post and Vote: ")
	} else {
		log.Println("[Post and Options] Block -> ", resp.BlockNum, " User -> ", author_name)
		return nil
	}
}

func (api *Client) Post_Options_Vote(author_name, title, body, permlink, ptag, post_image string, tags []string, weight_post int, percent uint16, votes, curation bool) error {
	if weight_post > 10000 {
		weight_post = 10000
	}
	if permlink == "" {
		permlink = translit.EncodeTitle(title)
	} else {
		permlink = translit.EncodeTitle(permlink)
	}
	tag := translit.EncodeTags(tags)
	if ptag == "" {
		ptag = translit.EncodeTag(tags[0])
	} else {
		ptag = translit.EncodeTag(ptag)
	}
	MAP := "1000000.000 GBG"
	PSD := percent
	if percent == 0 {
		MAP = "0.000 GBG"
	} else if percent == 50 {
		PSD = 10000
	} else {
		PSD = 0
	}

	json_meta := "{\"tags\":["
	for k, v := range tag {
		if k != len(tags)-1 {
			json_meta = json_meta + "\"" + v + "\","
		} else {
			json_meta = json_meta + "\"" + v + "\"]"
		}
	}
	if post_image != "" {
		json_meta = json_meta + ",\"image\":\"" + post_image + "\""
	}
	json_meta = json_meta + ",\"app\":\"golos-go(go-steem)\"}"

	var trx []types.Operation

	txp := &types.CommentOperation{
		ParentAuthor:   "",
		ParentPermlink: ptag,
		Author:         author_name,
		Permlink:       permlink,
		Title:          title,
		Body:           body,
		JsonMetadata:   json_meta,
	}
	trx = append(trx, txp)

	txo := &types.CommentOptionsOperation{
		Author:               author_name,
		Permlink:             permlink,
		MaxAcceptedPayout:    MAP,
		PercentSteemDollars:  PSD,
		AllowVotes:           votes,
		AllowCurationRewards: curation,
		Extensions:           []interface{}{},
	}
	trx = append(trx, txo)

	txv := &types.VoteOperation{
		Voter:    author_name,
		Author:   author_name,
		Permlink: permlink,
		Weight:   types.Int16(weight_post),
	}
	trx = append(trx, txv)

	resp, err := api.Send_Arr_Trx(author_name, trx)
	if err != nil {
		return errors.Wrapf(err, "Error Post and Vote: ")
	} else {
		log.Println("[Post and Options] Block -> ", resp.BlockNum, " User -> ", author_name)
		return nil
	}
}

func (api *Client) Follow(follower, following string) error {
	json_string := "[\"follow\",{\"follower\":\"" + follower + "\",\"following\":\"" + following + "\",\"what\":[\"blog\"]}]"

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{follower},
		ID:                   "follow",
		JSON:                 json_string,
	}
	resp, err := api.Send_Trx(follower, tx)
	if err != nil {
		return errors.Wrapf(err, "Error Reblog: ")
	} else {
		log.Println("[Follow] Block -> ", resp.BlockNum, " Follower user -> ", follower, " Following user -> ", following)
		return nil
	}
}

func (api *Client) Unfollow(follower, following string) error {
	json_string := "[\"follow\",{\"follower\":\"" + follower + "\",\"following\":\"" + following + "\",\"what\":[]}]"

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{follower},
		ID:                   "follow",
		JSON:                 json_string,
	}
	resp, err := api.Send_Trx(follower, tx)
	if err != nil {
		return errors.Wrapf(err, "Error Reblog: ")
	} else {
		log.Println("[Unfollow] Block -> ", resp.BlockNum, " Unfollower user -> ", follower, " Unfollowing user -> ", following)
		return nil
	}
}

func (api *Client) Ignore(follower, following string) error {
	json_string := "[\"follow\",{\"follower\":\"" + follower + "\",\"following\":\"" + following + "\",\"what\":[\"ignore\"]}]"

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{follower},
		ID:                   "follow",
		JSON:                 json_string,
	}
	resp, err := api.Send_Trx(follower, tx)
	if err != nil {
		return errors.Wrapf(err, "Error Reblog: ")
	} else {
		log.Println("[Ignore] Block -> ", resp.BlockNum, " Ignore user -> ", follower, " Ignoring user -> ", following)
		return nil
	}
}

func (api *Client) Notice(follower, following string) error {
	json_string := "[\"follow\",{\"follower\":\"" + follower + "\",\"following\":\"" + following + "\",\"what\":[]}]"

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{follower},
		ID:                   "follow",
		JSON:                 json_string,
	}
	resp, err := api.Send_Trx(follower, tx)
	if err != nil {
		return errors.Wrapf(err, "Error Reblog: ")
	} else {
		log.Println("[Notice] Block -> ", resp.BlockNum, " Notice user -> ", follower, " Noticing user -> ", following)
		return nil
	}
}

func (api *Client) Reblog(user_name, author_name, permlink string) error {
	if api.Verify_Reblogs(author_name, permlink, user_name) {
		return errors.New("The user already did repost")
	}
	json_string := "[\"reblog\",{\"account\":\"" + user_name + "\",\"author\":\"" + author_name + "\",\"permlink\":\"" + permlink + "\"]"

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{user_name},
		ID:                   "follow",
		JSON:                 json_string,
	}
	resp, err := api.Send_Trx(user_name, tx)
	if err != nil {
		return errors.Wrapf(err, "Error Reblog: ")
	} else {
		log.Println("[Reblog] Block -> ", resp.BlockNum, " Reblog user -> ", user_name, " Rebloging -> ", author_name, "/", permlink)
		return nil
	}
}

func (api *Client) AccountWitnessVote(user_name, witness_name string, approv bool) error {
	tx := &types.AccountWitnessVoteOperation{
		Account: user_name,
		Witness: witness_name,
		Approve: approv,
	}
	resp, err := api.Send_Trx(user_name, tx)
	if err != nil {
		return errors.Wrapf(err, "Error AccountWitnessVote: ")
	} else {
		log.Println("[AccountWitnessVote] Block -> ", resp.BlockNum, " User -> ", user_name)
		return nil
	}
}

func (api *Client) AccountWitnessProxy(user_name, proxy string) error {
	tx := &types.AccountWitnessProxyOperation{
		Account: user_name,
		Proxy:   proxy,
	}
	resp, err := api.Send_Trx(user_name, tx)
	if err != nil {
		return errors.Wrapf(err, "Error AccountWitnessProxy: ")
	} else {
		log.Println("[AccountWitnessProxy] Block -> ", resp.BlockNum, " User -> ", user_name)
		return nil
	}
}

func (api *Client) Transfer(from_name, to_name, memo, ammount string) error {
	tx := &types.TransferOperation{
		From:   from_name,
		To:     to_name,
		Amount: ammount,
		Memo:   memo,
	}
	resp, err := api.Send_Trx(from_name, tx)
	if err != nil {
		return errors.Wrapf(err, "Error Transfer: ")
	} else {
		log.Println("[Transfer] Block -> ", resp.BlockNum, " From user -> ", from_name, " To user -> ", to_name)
		return nil
	}
}

func (api *Client) Login(user_name, key string) bool {
	json_string := "[\"login\",{\"account\":\"" + user_name + "\",\"app\":\"golos-go(go-steem)\"}]"

	strx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{user_name},
		ID:                   "login",
		JSON:                 json_string,
	}

	props, err := api.Rpc.Database.GetDynamicGlobalProperties()
	if err != nil {
		return false
	}

	// Создание транзакции
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return false
	}
	tx := transactions.NewSignedTransaction(&types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	})

	// Добавление операций в транзакцию
	tx.PushOperation(strx)

	// Получаем необходимый для подписи ключ
	var keys [][]byte
	privKey, _ := wif.Decode(string([]byte(key)))
	keys = append(keys, privKey)

	// Подписываем транзакцию
	if err := tx.Sign(keys, api.Chain); err != nil {
		return false
	}

	// Отправка транзакции
	resp, err := api.Rpc.NetworkBroadcast.BroadcastTransactionSynchronous(tx.Transaction)

	if err != nil {
		return false
	} else {
		log.Println("[Login] Block -> ", resp.BlockNum, " User -> ", user_name)
		return true
	}
}

func (api *Client) LimitOrderCancel(owner string, orderid uint32) error {

	tx := &types.LimitOrderCancelOperation{
		Owner:   owner,
		OrderID: orderid,
	}

	resp, err := api.Send_Trx(owner, tx)
	if err != nil {
		return errors.Wrapf(err, "Error LimitOrderCancel: ")
	} else {
		log.Println("[LimitOrderCancel] Block -> ", resp.BlockNum, " LimitOrderCancel user -> ", owner)
		return nil
	}
}

func (api *Client) LimitOrderCreate(owner, sell, buy string, orderid uint32) error {

	expiration := time.Now().Add(3600000 * time.Second).UTC()
	fok := false

	tx := &types.LimitOrderCreateOperation{
		Owner:        owner,
		OrderID:      orderid,
		AmountToSell: sell,
		MinToReceive: buy,
		FillOrKill:   fok,
		Expiration:   &types.Time{&expiration},
	}

	resp, err := api.Send_Trx(owner, tx)
	if err != nil {
		return errors.Wrapf(err, "Error LimitOrderCreate: ")
	} else {
		log.Println("[LimitOrderCreate] Block -> ", resp.BlockNum, " LimitOrderCreate user -> ", owner)
		return nil
	}
}

func (api *Client) Convert(owner, amount string, requestid uint32) error {
	tx := &types.ConvertOperation{
		Owner:     owner,
		RequestID: requestid,
		Amount:    amount,
	}

	resp, err := api.Send_Trx(owner, tx)
	if err != nil {
		return errors.Wrapf(err, "Error Convert: ")
	} else {
		log.Println("[Convert] Block -> ", resp.BlockNum, " Convert user -> ", owner)
		return nil
	}
}

func (api *Client) TransferToVesting(from, to, amount string) error {
	tx := &types.TransferToVestingOperation{
		From:   from,
		To:     to,
		Amount: amount,
	}

	resp, err := api.Send_Trx(from, tx)
	if err != nil {
		return errors.Wrapf(err, "Error TransferToVesting: ")
	} else {
		log.Println("[TransferToVesting] Block -> ", resp.BlockNum, " TransferToVesting user -> ", from)
		return nil
	}
}

func (api *Client) WithdrawVesting(account, vshares string) error {
	tx := &types.WithdrawVestingOperation{
		Account:       account,
		VestingShares: vshares,
	}

	resp, err := api.Send_Trx(account, tx)
	if err != nil {
		return errors.Wrapf(err, "Error WithdrawVesting: ")
	} else {
		log.Println("[WithdrawVesting] Block -> ", resp.BlockNum, " WithdrawVesting user -> ", account)
		return nil
	}
}

func (api *Client) ChangeRecoveryAccount(accounttorecover, newrecoveryaccount string) error {
	tx := &types.ChangeRecoveryAccountOperation{
		AccountToRecover:   accounttorecover,
		NewRecoveryAccount: newrecoveryaccount,
		Extensions:         []interface{}{},
	}

	resp, err := api.Send_Trx(accounttorecover, tx)
	if err != nil {
		return errors.Wrapf(err, "Error ChangeRecoveryAccount: ")
	} else {
		log.Println("[ChangeRecoveryAccount] Block -> ", resp.BlockNum, " ChangeRecoveryAccount user -> ", accounttorecover)
		return nil
	}
}

func (api *Client) TransferToSavings(from, to, amount, memo string) error {
	tx := &types.TransferToSavingsOperation{
		From:   from,
		To:     to,
		Amount: amount,
		Memo:   memo,
	}

	resp, err := api.Send_Trx(from, tx)
	if err != nil {
		return errors.Wrapf(err, "Error TransferToSavings: ")
	} else {
		log.Println("[TransferToSavings] Block -> ", resp.BlockNum, " TransferToSavings user -> ", from)
		return nil
	}
}

func (api *Client) TransferFromSavings(from, to, amount, memo string, requestid uint32) error {
	tx := &types.TransferFromSavingsOperation{
		From:      from,
		RequestId: requestid,
		To:        to,
		Amount:    amount,
		Memo:      memo,
	}

	resp, err := api.Send_Trx(from, tx)
	if err != nil {
		return errors.Wrapf(err, "Error TransferFromSavings: ")
	} else {
		log.Println("[TransferFromSavings] Block -> ", resp.BlockNum, " TransferFromSavings user -> ", from)
		return nil
	}
}

func (api *Client) CancelTransferFromSavings(from string, requestid uint32) error {
	tx := &types.CancelTransferFromSavingsOperation{
		From:      from,
		RequestId: requestid,
	}

	resp, err := api.Send_Trx(from, tx)
	if err != nil {
		return errors.Wrapf(err, "Error CancelTransferFromSavings: ")
	} else {
		log.Println("[CancelTransferFromSavings] Block -> ", resp.BlockNum, " CancelTransferFromSavings user -> ", from)
		return nil
	}
}

func (api *Client) DeclineVotingRights(account string, decline bool) error {
	tx := &types.DeclineVotingRightsOperation{
		Account: account,
		Decline: decline,
	}

	resp, err := api.Send_Trx(account, tx)
	if err != nil {
		return errors.Wrapf(err, "Error DeclineVotingRights: ")
	} else {
		log.Println("[DeclineVotingRights] Block -> ", resp.BlockNum, " DeclineVotingRights user -> ", account)
		return nil
	}
}

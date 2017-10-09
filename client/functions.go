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

func (api *Client) Post(author_name, title, body, permlink, ptag string, tags []string) error {
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
			json_meta = json_meta + "\"" + v + "\"],\"app\":\"golos-go(go-steem)\"}"
		}
	}

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

func (api *Client) Post_Vote(author_name, title, body, permlink, ptag string, tags []string, weight_post int) error {
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
			json_meta = json_meta + "\"" + v + "\"],\"app\":\"golos-go(go-steem)\"}"
		}
	}
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

func (api *Client) Post_Options(author_name, title, body, permlink, ptag string, tags []string, percent uint16, votes, curation bool) error {
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
			json_meta = json_meta + "\"" + v + "\"],\"app\":\"golos-go(go-steem)\"}"
		}
	}
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

func (api *Client) Post_Options_Vote(author_name, title, body, permlink, ptag string, tags []string, weight_post int, percent uint16, votes, curation bool) error {
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
			json_meta = json_meta + "\"" + v + "\"],\"app\":\"golos-go(go-steem)\"}"
		}
	}
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

func (api *Client) Witness_Vote(user_name, witness_name string, approv bool) error {
	tx := &types.AccountWitnessVoteOperation{
		Account: user_name,
		Witness: witness_name,
		Approve: approv,
	}
	resp, err := api.Send_Trx(user_name, tx)
	if err != nil {
		return errors.Wrapf(err, "Error Reblog: ")
	} else {
		log.Println("[Witness Vote] Block -> ", resp.BlockNum, " User -> ", user_name, " Witness user -> ", witness_name)
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

func (api *Client) FeedPublish(username, base, quote string) error {

	ExchR := types.ExchRate{Base: base, Quote: quote}
	tx := &types.FeedPublishOperation{
		Publisher:    username,
		ExchangeRate: ExchR,
	}

	resp, err := api.Send_Trx(username, tx)
	if err != nil {
		return errors.Wrapf(err, "Error FeedPublish: ")
	} else {
		log.Println("[FeedPublish] Block -> ", resp.BlockNum, " FeedPublish user -> ", username)
		return nil
	}
}

func (api *Client) WitnessUpdate(username, url, signKey, acfee, fee string, blocksize uint32, sbdir uint16) error {

	chprop := types.ChainProperties{AccountCreationFee: acfee, MaximumBlockSize: blocksize, SBDInterestRate: sbdir}
	tx := &types.WitnessUpdateOperation{
		Owner:           username,
		Url:             url,
		BlockSigningKey: signKey,
		Props:           &chprop,
		Fee:             fee,
	}

	resp, err := api.Send_Trx(username, tx)
	if err != nil {
		return errors.Wrapf(err, "Error WitnessUpdate: ")
	} else {
		log.Println("[WitnessUpdate] Block -> ", resp.BlockNum, " WitnessUpdate user -> ", username)
		return nil
	}
}

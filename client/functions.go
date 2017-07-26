package client

import (
	// Stdlib
	"log"
	"strconv"
	"time"

	// Vendor
	"github.com/pkg/errors"

	// RPC
	"github.com/asuleymanov/golos-go/translit"
	"github.com/asuleymanov/golos-go/types"
)

func (api *Client) Vote(user_name, author_name, permlink string, weight int) error {
	if weight > 10000 {
		return errors.New("The value of Weight can not be more than 10,000")
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
	resp, err := api.Send_Trx(tx)
	if err != nil {
		return errors.Wrapf(err, "Error Vote: ")
	} else {
		log.Println("Add Vote to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
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
	resp, err := api.Send_Trx(tx)
	if err != nil {
		return errors.Wrapf(err, "Error Comment: ")
	} else {
		log.Println("Add Comment to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
		return nil
	}
}

func (api *Client) Comment_Vote(user_name, author_name, ppermlink, body string, weight_post int) error {
	if weight_post > 10000 {
		return errors.New("The value of Weight can not be more than 10,000")
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

	resp, err := api.Send_Arr_Trx(trx)
	if err != nil {
		return errors.Wrapf(err, "Error Comment and Vote: ")
	} else {
		log.Println("Add Comment and Vote to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
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
	resp, err := api.Send_Trx(tx)
	if err != nil {
		return errors.Wrapf(err, "Error Delete Comment: ")
	} else {
		log.Println("Delete Comment to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
		return nil
	}
}

func (api *Client) Post(author_name, title, body string, tags []string) error {
	permlink := translit.EncodeTitle(title)
	tag := translit.EncodeTags(tags)
	ptag := translit.EncodeTag(tags[0])

	json_meta := "{\"tag\":["
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

	resp, err := api.Send_Trx(tx)
	if err != nil {
		return errors.Wrapf(err, "Error Post: ")
	} else {
		log.Println("Add Post to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
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
	resp, err := api.Send_Trx(tx)
	if err != nil {
		return errors.Wrapf(err, "Error Reblog: ")
	} else {
		log.Println("Follow to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
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
	resp, err := api.Send_Trx(tx)
	if err != nil {
		return errors.Wrapf(err, "Error Reblog: ")
	} else {
		log.Println("Unfollow to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
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
	resp, err := api.Send_Trx(tx)
	if err != nil {
		return errors.Wrapf(err, "Error Reblog: ")
	} else {
		log.Println("Ignore to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
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
	resp, err := api.Send_Trx(tx)
	if err != nil {
		return errors.Wrapf(err, "Error Reblog: ")
	} else {
		log.Println("Notice to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
		return nil
	}
}

func (api *Client) Reblog(user_name, author_name, permlink string) error {
	if api.Verify_Reblogs(author_name, permlink, user_name) {
		return errors.New("The user already did repost")
	}
	json_string := "[\"reblog\",{\"account\":\"" + user_name + "\",\"author\":\"" + author_name + "\",\"permlink\":\"" + permlink + "\"}]"

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{user_name},
		ID:                   "follow",
		JSON:                 json_string,
	}
	resp, err := api.Send_Trx(tx)
	if err != nil {
		return errors.Wrapf(err, "Error Reblog: ")
	} else {
		log.Println("Reblog to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
		return nil
	}
}

func (api *Client) Witness_Vote(user_name, witness_name string, approv bool) error {
	tx := &types.AccountWitnessVoteOperation{
		Account: user_name,
		Witness: witness_name,
		Approve: approv,
	}
	resp, err := api.Send_Trx(tx)
	if err != nil {
		return errors.Wrapf(err, "Error Reblog: ")
	} else {
		log.Println("Witness Vote to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
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
	resp, err := api.Send_Trx(tx)
	if err != nil {
		return errors.Wrapf(err, "Error Reblog: ")
	} else {
		log.Println("Transfer to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
		return nil
	}
}

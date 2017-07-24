package client

import (
	// Stdlib
	"log"
	"strconv"
	"time"

	// Vendor
	"github.com/pkg/errors"

	// RPC
	_ "github.com/asuleymanov/golos-go/translit"
	"github.com/asuleymanov/golos-go/types"
)

func (api *Client) Vote(author, permlink string, weight int) error {
	if weight > 10000 {
		return errors.New("The value of Weight can not be more than 10,000")
	}
	if api.Verify_Voter(author, permlink, api.User.Name) {
		return errors.New("The voter is on the list")
	}
	tx := &types.VoteOperation{
		Voter:    api.User.Name,
		Author:   author,
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

func (api *Client) Comment(author, ppermlink, body string) error {
	times, _ := strconv.Unquote(time.Now().Add(30 * time.Second).UTC().Format(fdt))
	permlink := "re-" + author + "-" + ppermlink + "-" + times
	tx := &types.CommentOperation{
		ParentAuthor:   author,
		ParentPermlink: ppermlink,
		Author:         api.User.Name,
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

func (api *Client) Comment_Vote(author, ppermlink, body string, weight_post int) error {
	times, _ := strconv.Unquote(time.Now().Add(30 * time.Second).UTC().Format(fdt))
	permlink := "re-" + author + "-" + ppermlink + "-" + times
	var trx []types.Operation
	txc := &types.CommentOperation{
		ParentAuthor:   author,
		ParentPermlink: ppermlink,
		Author:         api.User.Name,
		Permlink:       permlink,
		Title:          "",
		Body:           body,
		JsonMetadata:   "{\"app\":\"golos-go(go-steem)\"}",
	}
	trx = append(trx, txc)

	if !api.Verify_Voter(author, permlink, api.User.Name) {
		txv := &types.VoteOperation{
			Voter:    api.User.Name,
			Author:   author,
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

func (api *Client) DeleteComment(permlink string) error {
	if api.Verify_Votes(api.User.Name, permlink) {
		return errors.New("You can not delete already there are voted")
	}
	if api.Verify_Comments(api.User.Name, permlink) {
		return errors.New("You can not delete already have comments")
	}
	tx := &types.DeleteCommentOperation{
		Author:   api.User.Name,
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

/*
func (api *Client) Post(title, body string, tags []string) error {
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
		Author:         api.User.Name,
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

func (api *Client) Follow(author string) error {
	json_string := "[\"follow\",{\"follower\":\"" + api.User.Name + "\",\"following\":\"" + author + "\",\"what\":[\"blog\"]}]"

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{api.User.Name},
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

func (api *Client) Reblog(author, permlink string) error {
	json_string := "[\"reblog\",{\"account\":\"" + api.User.Name + "\",\"author\":\"" + author + "\",\"permlink\":\"" + permlink + "\"}]"

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{api.User.Name},
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

func (api *Client) Unfollow(author string) error {
	json_string := "[\"follow\",{\"follower\":\"" + api.User.Name + "\",\"following\":\"" + author + "\",\"what\":[\"\"]}]"

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{api.User.Name},
		ID:                   "reblog",
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

func (api *Client) Ignore(author string) error {
	json_string := "[\"follow\",{\"follower\":\"" + api.User.Name + "\",\"following\":\"" + author + "\",\"what\":[\"ignore\"]}]"

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{api.User.Name},
		ID:                   "reblog",
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
*/

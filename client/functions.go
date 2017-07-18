package client

import (
	// Stdlib
	"log"
	"strconv"
	"time"

	// Vendor
	"github.com/pkg/errors"

	// RPC
	"github.com/asuleymanov/golos-go/types"
)

func (api *Golos) Vote(author, permlink string, weight int, chain string) error {
	if weight > 10000 {
		return errors.New("The value of Weight can not be more than 10,000")
	}
	tx := &types.VoteOperation{
		Voter:    api.User.Name,
		Author:   author,
		Permlink: permlink,
		Weight:   types.Int16(weight),
	}
	resp, err := api.Send_Trx(tx, chain)
	if err != nil {
		return errors.Wrapf(err, "Error Vote: ")
	} else {
		log.Println("Add Vote to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
		return nil
	}
}

func (api *Golos) Comment(author, ppermlink, body, chain string) error {
	times, _ := strconv.Unquote(time.Now().Add(30 * time.Second).UTC().Format(fdt))
	permlink := "re-" + author + "-" + ppermlink + "-" + times
	tx := &types.CommentOperation{
		ParentAuthor:   author,
		ParentPermlink: ppermlink,
		Author:         api.User.Name,
		Permlink:       permlink,
		Title:          "",
		Body:           body,
		JsonMetadata:   "",
	}
	resp, err := api.Send_Trx(tx, chain)
	if err != nil {
		return errors.Wrapf(err, "Error Comment: ")
	} else {
		log.Println("Add Comment to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
		return nil
	}
}

func (api *Golos) CommentAndVote(author, ppermlink, body string, weight int, chain string) error {
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
		JsonMetadata:   "",
	}
	trx = append(trx, txc)

	txv := &types.VoteOperation{
		Voter:    api.User.Name,
		Author:   author,
		Permlink: ppermlink,
		Weight:   types.Int16(weight),
	}
	trx = append(trx, txv)

	resp, err := api.Send_Arr_Trx(trx, chain)
	if err != nil {
		return errors.Wrapf(err, "Error Comment and Vote: ")
	} else {
		log.Println("Add Comment and Vote to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
		return nil
	}
}

func (api *Golos) DeleteComment(permlink, chain string) error {
	tx := &types.DeleteCommentOperation{
		Author:   api.User.Name,
		Permlink: permlink,
	}
	resp, err := api.Send_Trx(tx, chain)
	if err != nil {
		return errors.Wrapf(err, "Error Delete Comment: ")
	} else {
		log.Println("Delete Comment to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
		return nil
	}
}

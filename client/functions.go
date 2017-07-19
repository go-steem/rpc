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

func (api *Golos) Vote(author, permlink string, weight int) error {
	if weight > 10000 {
		return errors.New("The value of Weight can not be more than 10,000")
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

func (api *Golos) Comment(author, ppermlink, body string) error {
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

func (api *Golos) Comment_Vote(author, ppermlink, body string, weight_post int) error {
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

	txv := &types.VoteOperation{
		Voter:    api.User.Name,
		Author:   author,
		Permlink: ppermlink,
		Weight:   types.Int16(weight_post),
	}
	trx = append(trx, txv)

	resp, err := api.Send_Arr_Trx(trx)
	if err != nil {
		return errors.Wrapf(err, "Error Comment and Vote: ")
	} else {
		log.Println("Add Comment and Vote to Block -> ", resp.BlockNum, " Trx -> ", resp.ID)
		return nil
	}
}

func (api *Golos) Post(title, body string, tags []string) error {
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
		ParentPermlink: permlink,
		Author:         api.User.Name,
		Permlink:       ptag,
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

func (api *Golos) DeleteComment(permlink string) error {
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

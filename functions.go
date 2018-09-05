package client

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/asuleymanov/steem-go/encoding/wif"
	"github.com/asuleymanov/steem-go/transactions"
	"github.com/asuleymanov/steem-go/translit"
	"github.com/asuleymanov/steem-go/types"
)

const fdt = `"20060102t150405"`

//Vote for publication
func (client *Client) Vote(username, authorname, permlink string, weight int) (*OperResp, error) {
	if weight > 10000 {
		weight = 10000
	}
	if client.VerifyVoterWeight(authorname, permlink, username, weight) {
		return nil, errors.New("The voter is on the list")
	}

	var trx []types.Operation

	tx := &types.VoteOperation{
		Voter:    username,
		Author:   authorname,
		Permlink: permlink,
		Weight:   types.Int16(weight),
	}
	trx = append(trx, tx)

	resp, err := client.SendTrx(username, trx)
	return &OperResp{NameOper: "Vote", Bresp: resp}, err
}

//MultiVote mass voting for publication.
//Using the opportunity to delegate the rights to sign.
func (client *Client) MultiVote(username, author, permlink string, arrvote []ArrVote) (*OperResp, error) {
	var trx []types.Operation
	var arrvotes []ArrVote

	for _, v := range arrvote {
		if client.VerifyDelegatePostingKeySign(v.User, username) && !client.VerifyVoter(author, permlink, v.User) {
			arrvotes = append(arrvotes, v)
		}
	}

	if len(arrvotes) == 0 {
		return nil, errors.New("Error Multi_Vote : All users from the list have already voted")
	}

	for _, val := range arrvotes {
		txt := &types.VoteOperation{
			Voter:    val.User,
			Author:   author,
			Permlink: permlink,
			Weight:   types.Int16(val.Weight),
		}
		trx = append(trx, txt)
	}

	resp, err := client.SendTrx(username, trx)
	return &OperResp{NameOper: "MultiVote", Bresp: resp}, err
}

//Comment for publication
func (client *Client) Comment(username, authorname, ppermlink, body string, o *PCOptions) (*OperResp, error) {
	var trx []types.Operation

	times, _ := strconv.Unquote(time.Now().Add(30 * time.Second).UTC().Format(fdt))
	permlink := "re-" + authorname + "-" + ppermlink + "-" + times
	permlink = strings.Replace(permlink, ".", "-", -1)

	jsonMeta := &types.ContentMetadata{"lib": "steem-go"}

	tx := &types.CommentOperation{
		ParentAuthor:   authorname,
		ParentPermlink: ppermlink,
		Author:         username,
		Permlink:       permlink,
		Title:          "",
		Body:           body,
		JSONMetadata:   jsonMeta,
	}
	trx = append(trx, tx)

	if o != nil {
		trx = append(trx, GetCommentOptionsOperation(username, permlink, *o))
	}

	resp, err := client.SendTrx(username, trx)
	return &OperResp{NameOper: "Comment", PermLink: permlink, Bresp: resp}, err
}

//Post creating a publication
func (client *Client) Post(authorname, title, body, permlink, ptag, postImage string, tags []string, o *PCOptions) (*OperResp, error) {
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

	jsonMeta := &types.ContentMetadata{
		"tags":  tag,
		"image": []string{postImage},
		"lib":   "steem-go",
	}

	var trx []types.Operation
	txp := &types.CommentOperation{
		ParentAuthor:   "",
		ParentPermlink: ptag,
		Author:         authorname,
		Permlink:       permlink,
		Title:          title,
		Body:           body,
		JSONMetadata:   jsonMeta,
	}
	trx = append(trx, txp)

	if o != nil {
		trx = append(trx, GetCommentOptionsOperation(authorname, permlink, *o))
	}

	resp, err := client.SendTrx(authorname, trx)
	return &OperResp{NameOper: "Post", PermLink: permlink, Bresp: resp}, err
}

//DeleteComment deleting a publication or comment
func (client *Client) DeleteComment(authorname, permlink string) (*OperResp, error) {
	if client.VerifyVotes(authorname, permlink) {
		return nil, errors.New("You can not delete already there are voted")
	}
	if client.VerifyComments(authorname, permlink) {
		return nil, errors.New("You can not delete already have comments")
	}
	var trx []types.Operation

	tx := &types.DeleteCommentOperation{
		Author:   authorname,
		Permlink: permlink,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(authorname, trx)
	return &OperResp{NameOper: "Delete Comment/Post", Bresp: resp}, err
}

//Follows subscribe(unsubscribe,ignore) to the user
/*
what:
blog
ignore
empty value
*/
func (client *Client) Follows(follower, following, what string) (*OperResp, error) {
	var trx []types.Operation
	js := types.FollowOperation{
		Follower:  follower,
		Following: following,
		What:      []string{what},
	}

	jsonString, errj := types.MarshalCustomJSON(js)
	if errj != nil {
		return nil, errj
	}

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{follower},
		ID:                   "follow",
		JSON:                 jsonString,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(follower, trx)
	respOper := ""
	switch what {
	case "":
		respOper = "Neutrality"
	case "blog":
		respOper = "Follows"
	case "ignore":
		respOper = "Ignore"
	}
	return &OperResp{NameOper: respOper, Bresp: resp}, err
}

//Reblog repost records
func (client *Client) Reblog(username, authorname, permlink string) (*OperResp, error) {
	if client.VerifyReblogs(authorname, permlink, username) {
		return nil, errors.New("The user already did repost")
	}

	js := types.ReblogOperation{
		Account:  username,
		Author:   authorname,
		Permlink: permlink,
	}

	jsonString, errj := types.MarshalCustomJSON(js)
	if errj != nil {
		return nil, errj
	}

	var trx []types.Operation

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{username},
		ID:                   "follow",
		JSON:                 jsonString,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(username, trx)
	return &OperResp{NameOper: "Reblog", Bresp: resp}, err
}

//AccountWitnessVote of voting for the delegate.
func (client *Client) AccountWitnessVote(username, witnessName string, approv bool) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.AccountWitnessVoteOperation{
		Account: username,
		Witness: witnessName,
		Approve: approv,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(username, trx)
	return &OperResp{NameOper: "AccountWitnessVote", Bresp: resp}, err
}

//AccountWitnessProxy transfer of the right to vote for delegates to another user.
func (client *Client) AccountWitnessProxy(username, proxy string) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.AccountWitnessProxyOperation{
		Account: username,
		Proxy:   proxy,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(username, trx)
	return &OperResp{NameOper: "AccountWitnessProxy", Bresp: resp}, err
}

//Transfer of funds to any user.
func (client *Client) Transfer(fromName, toName, memo, ammount string) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.TransferOperation{
		From:   fromName,
		To:     toName,
		Amount: ammount,
		Memo:   memo,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(fromName, trx)
	return &OperResp{NameOper: "Transfer", Bresp: resp}, err
}

//MultiTransfer multiple funds transfer in one transaction.
func (client *Client) MultiTransfer(username string, arrtrans []ArrTransfer) (*OperResp, error) {
	var trx []types.Operation

	for _, val := range arrtrans {
		txt := &types.TransferOperation{
			From:   username,
			To:     val.To,
			Amount: val.Ammount,
			Memo:   val.Memo,
		}
		trx = append(trx, txt)
	}

	resp, err := client.SendTrx(username, trx)
	return &OperResp{NameOper: "MultiTransfer", Bresp: resp}, err
}

//Login checking the user's key for the possibility of operations in GOLOS.
func (client *Client) Login(username, key string) bool {
	jsonString := "[\"login\",{\"account\":\"" + username + "\",\"app\":\"steem-go\"}]"

	strx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{username},
		ID:                   "login",
		JSON:                 jsonString,
	}

	props, err := client.Database.GetDynamicGlobalProperties()
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
	if err := tx.Sign(keys, client.Chain); err != nil {
		return false
	}

	// Отправка транзакции
	resp, err := client.NetworkBroadcast.BroadcastTransactionSynchronous(tx.Transaction)

	if err != nil {
		return false
	}
	log.Println("[Login] Block -> ", resp.BlockNum, " User -> ", username)
	return true
}

//LimitOrderCancel restrict order Cancel
func (client *Client) LimitOrderCancel(owner string, orderid uint32) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.LimitOrderCancelOperation{
		Owner:   owner,
		OrderID: orderid,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(owner, trx)
	return &OperResp{NameOper: "LimitOrderCancel", Bresp: resp}, err
}

//LimitOrderCreate Creating a limit order
func (client *Client) LimitOrderCreate(owner, sell, buy string, orderid uint32) (*OperResp, error) {
	var trx []types.Operation

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

	trx = append(trx, tx)
	resp, err := client.SendTrx(owner, trx)
	return &OperResp{NameOper: "LimitOrderCreate", Bresp: resp}, err
}

//Convert conversion
func (client *Client) Convert(owner, amount string, requestid uint32) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.ConvertOperation{
		Owner:     owner,
		RequestID: requestid,
		Amount:    amount,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(owner, trx)
	return &OperResp{NameOper: "Convert", Bresp: resp}, err
}

//TransferToVesting transfer to POWER
func (client *Client) TransferToVesting(from, to, amount string) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.TransferToVestingOperation{
		From:   from,
		To:     to,
		Amount: amount,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(from, trx)
	return &OperResp{NameOper: "TransferToVesting", Bresp: resp}, err
}

//WithdrawVesting down POWER
func (client *Client) WithdrawVesting(account, vshares string) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.WithdrawVestingOperation{
		Account:       account,
		VestingShares: vshares,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(account, trx)
	return &OperResp{NameOper: "WithdrawVesting", Bresp: resp}, err
}

//ChangeRecoveryAccount change account with which you can restore access
func (client *Client) ChangeRecoveryAccount(accounttorecover, newrecoveryaccount string) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.ChangeRecoveryAccountOperation{
		AccountToRecover:   accounttorecover,
		NewRecoveryAccount: newrecoveryaccount,
		Extensions:         []interface{}{},
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(accounttorecover, trx)
	return &OperResp{NameOper: "ChangeRecoveryAccount", Bresp: resp}, err
}

//TransferToSavings transfer to safe
func (client *Client) TransferToSavings(from, to, amount, memo string) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.TransferToSavingsOperation{
		From:   from,
		To:     to,
		Amount: amount,
		Memo:   memo,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(from, trx)
	return &OperResp{NameOper: "TransferToSavings", Bresp: resp}, err
}

//TransferFromSavings transfer from safe
func (client *Client) TransferFromSavings(from, to, amount, memo string, requestid uint32) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.TransferFromSavingsOperation{
		From:      from,
		RequestID: requestid,
		To:        to,
		Amount:    amount,
		Memo:      memo,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(from, trx)
	return &OperResp{NameOper: "TransferFromSavings", Bresp: resp}, err
}

//CancelTransferFromSavings cancel transfer to safe
func (client *Client) CancelTransferFromSavings(from string, requestid uint32) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.CancelTransferFromSavingsOperation{
		From:      from,
		RequestID: requestid,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(from, trx)
	return &OperResp{NameOper: "CancelTransferFromSavings", Bresp: resp}, err
}

//DeclineVotingRights disabling the possibility of any vote.
//It is impossible to restore the possibility of voting
func (client *Client) DeclineVotingRights(account string, decline bool) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.DeclineVotingRightsOperation{
		Account: account,
		Decline: decline,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(account, trx)
	return &OperResp{NameOper: "DeclineVotingRights", Bresp: resp}, err
}

//FeedPublish update course data
func (client *Client) FeedPublish(publisher, base, quote string) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.FeedPublishOperation{
		Publisher: publisher,
		ExchangeRate: &types.ExchRate{
			Base:  base,
			Quote: quote,
		},
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(publisher, trx)
	return &OperResp{NameOper: "FeedPublish", Bresp: resp}, err
}

//WitnessUpdate updating delegate data
func (client *Client) WitnessUpdate(owner, url, blocksigningkey, accountcreationfee string, maxblocksize uint32, sbdinterestrate uint16) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.WitnessUpdateOperation{
		Owner:           owner,
		URL:             url,
		BlockSigningKey: blocksigningkey,
		Props: &types.ChainProperties{
			AccountCreationFee: accountcreationfee,
			MaximumBlockSize:   maxblocksize,
			SBDInterestRate:    sbdinterestrate,
		},
		Fee: "0.000 GOLOS",
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(owner, trx)
	return &OperResp{NameOper: "WitnessUpdate", Bresp: resp}, err
}

//AccountCreate creating a user in systems
func (client *Client) AccountCreate(creator, newAccountName, password, fee string) (*OperResp, error) {
	type Keys struct {
		Private string
		Public  string
	}

	var trx []types.Operation
	var listKeys = make(map[string]Keys)

	empty := map[string]int64{}
	roles := [4]string{"owner", "active", "posting", "memo"}

	for _, val := range roles {
		priv := GetPrivateKey(newAccountName, val, password)
		pub := GetPublicKey("GLS", priv)
		listKeys[val] = Keys{Private: priv, Public: pub}
	}

	owner := types.Authority{
		WeightThreshold: 1,
		AccountAuths:    empty,
		KeyAuths:        map[string]int64{listKeys["owner"].Public: 1},
	}

	active := types.Authority{
		WeightThreshold: 1,
		AccountAuths:    empty,
		KeyAuths:        map[string]int64{listKeys["active"].Public: 1},
	}

	posting := types.Authority{
		WeightThreshold: 1,
		AccountAuths:    empty,
		KeyAuths:        map[string]int64{listKeys["posting"].Public: 1},
	}

	jsonMeta := &types.AccountMetadata{}

	tx := &types.AccountCreateOperation{
		Fee:            fee,
		Creator:        creator,
		NewAccountName: newAccountName,
		Owner:          &owner,
		Active:         &active,
		Posting:        &posting,
		MemoKey:        listKeys["memo"].Public,
		JSONMetadata:   jsonMeta,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(creator, trx)
	return &OperResp{NameOper: "AccountCreate", Bresp: resp}, err
}

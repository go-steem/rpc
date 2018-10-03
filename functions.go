package client

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/asuleymanov/steem-go/types"
	"github.com/asuleymanov/translit"
)

const fdt = `"20060102t150405"`

//Vote for publication
func (client *Client) Vote(username, authorname, permlink string, weight int) (*OperResp, error) {
	if weight > 10000 {
		return nil, errors.New("The weight can not be greater than 10,000 and less than -10000")
	}
	bvvw, evvw := client.VerifyVoterWeight(authorname, permlink, username, weight)
	if evvw != nil {
		return nil, evvw
	} else if evvw == nil && !bvvw {
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
		bvdpks, evdpks := client.VerifyDelegatePostingKeySign(v.User, username)
		bvv, evv := client.VerifyVoter(author, permlink, v.User)
		if evdpks != nil {
			return nil, evdpks
		}
		if evv != nil {
			return nil, evv
		}
		if bvdpks && !bvv {
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

	times, errUnq := strconv.Unquote(time.Now().Add(30 * time.Second).UTC().Format(fdt))
	if errUnq != nil {
		return nil, errUnq
	}

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
		JSONMetadata:   client.GenCommentMetadata(jsonMeta),
	}
	trx = append(trx, tx)

	if o != nil {
		trx = append(trx, GetCommentOptionsOperation(username, permlink, *o))
	}

	resp, err := client.SendTrx(username, trx)
	return &OperResp{NameOper: "Comment", PermLink: permlink, Bresp: resp}, err
}

//Post creating a publication
func (client *Client) Post(authorname, title, body, permlink, ppermlink, postImage string, tags []string, o *PCOptions) (*OperResp, error) {
	if permlink == "" {
		var err error
		permlink, err = translit.EncodeTitle(title)
		if err != nil {
			return nil, err
		}
	}

	tag := translit.EncodeTags(tags)
	if ppermlink == "" {
		ppermlink = translit.EncodeTag(tags[0])
	} else {
		ppermlink = translit.EncodeTag(ppermlink)
	}

	jsonMeta := &types.ContentMetadata{
		"tags":  tag,
		"image": []string{postImage},
		"lib":   "steem-go",
	}

	var trx []types.Operation
	txp := &types.CommentOperation{
		ParentAuthor:   "",
		ParentPermlink: ppermlink,
		Author:         authorname,
		Permlink:       permlink,
		Title:          title,
		Body:           body,
		JSONMetadata:   client.GenCommentMetadata(jsonMeta),
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
	bvv, evv := client.VerifyVotes(authorname, permlink)
	if evv != nil {
		return nil, evv
	} else if evv == nil && !bvv {
		return nil, errors.New("You can not delete already there are voted")
	}

	bec, eec := client.ExistComments(authorname, permlink)
	if eec != nil {
		return nil, eec
	} else if eec == nil && !bec {
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
	bvr, evr := client.VerifyReblogs(authorname, permlink, username)
	if evr != nil {
		return nil, evr
	} else if evr == nil && !bvr {
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
func (client *Client) Transfer(fromName, toName, memo string, amount *types.Asset) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.TransferOperation{
		From:   fromName,
		To:     toName,
		Amount: amount,
		Memo:   memo,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(fromName, trx)
	return &OperResp{NameOper: "Transfer", Bresp: resp}, err
}

//MultiTransfer multiple funds transfer in one transaction.
func (client *Client) MultiTransfer(username string, arrtrans []ArrTransfer) (*OperResp, error) {
	var trx []types.Operation

	for i := range arrtrans {
		txt := &types.TransferOperation{
			From:   username,
			To:     arrtrans[i].To,
			Amount: &arrtrans[i].Amount,
			Memo:   arrtrans[i].Memo,
		}
		trx = append(trx, txt)
	}

	resp, err := client.SendTrx(username, trx)
	return &OperResp{NameOper: "MultiTransfer", Bresp: resp}, err
}

//Login checking the user's posting key for the possibility of operations in STEEM.
func (client *Client) Login(username string) (bool, error) {
	js := types.LoginOperation{
		Account: username,
	}

	jsonString, errj := types.MarshalCustomJSON(js)
	if errj != nil {
		return false, errj
	}

	var trx []types.Operation

	tx := &types.CustomJSONOperation{
		RequiredAuths:        []string{},
		RequiredPostingAuths: []string{username},
		ID:                   "login",
		JSON:                 jsonString,
	}

	trx = append(trx, tx)
	_, err := client.SendTrx(username, trx)
	if err != nil {
		return false, err
	}

	return true, nil
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

//LimitOrderCreate сreating a limit order
func (client *Client) LimitOrderCreate(owner string, sell, buy *types.Asset, orderid uint32) (*OperResp, error) {
	var trx []types.Operation

	expiration := time.Now().Add(3600000 * time.Second).UTC()

	tx := &types.LimitOrderCreateOperation{
		Owner:        owner,
		OrderID:      orderid,
		AmountToSell: sell,
		MinToReceive: buy,
		FillOrKill:   false,
		Expiration:   &types.Time{Time: &expiration},
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(owner, trx)
	return &OperResp{NameOper: "LimitOrderCreate", Bresp: resp}, err
}

//LimitOrderCreate2 сreation of a limit order based on a certain rate.
func (client *Client) LimitOrderCreate2(owner string, sell, base, quote *types.Asset, orderid uint32) (*OperResp, error) {
	var trx []types.Operation

	expiration := time.Now().Add(3600000 * time.Second).UTC()

	tx := &types.LimitOrderCreate2Operation{
		Owner:        owner,
		OrderID:      orderid,
		AmountToSell: sell,
		ExchangeRate: &types.ExchRate{
			Base:  base,
			Quote: quote,
		},
		FillOrKill: false,
		Expiration: &types.Time{Time: &expiration},
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(owner, trx)
	return &OperResp{NameOper: "LimitOrderCreate2", Bresp: resp}, err
}

//Convert conversion
func (client *Client) Convert(owner string, amount *types.Asset, requestid uint32) (*OperResp, error) {
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
func (client *Client) TransferToVesting(from, to string, amount *types.Asset) (*OperResp, error) {
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
func (client *Client) WithdrawVesting(account string, vshares *types.Asset) (*OperResp, error) {
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
func (client *Client) TransferToSavings(from, to, memo string, amount *types.Asset) (*OperResp, error) {
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
func (client *Client) TransferFromSavings(from, to, memo string, amount *types.Asset, requestid uint32) (*OperResp, error) {
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
func (client *Client) FeedPublish(publisher string, base, quote *types.Asset) (*OperResp, error) {
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
func (client *Client) WitnessUpdate(owner, url, blocksigningkey string) (*OperResp, error) {
	var trx []types.Operation

	if url == "" {
		url = " "
	}

	ans, erra := client.Database.GetWitnessByAccount(owner)
	if erra != nil {
		return nil, erra
	}

	tx := &types.WitnessUpdateOperation{
		Owner:           owner,
		URL:             url,
		BlockSigningKey: blocksigningkey,
		Props:           ans.Props,
		Fee:             SetAsset(0.000, "STEEM"),
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(owner, trx)
	return &OperResp{NameOper: "WitnessUpdate", Bresp: resp}, err
}

//AccountCreate creating a user in systems
func (client *Client) AccountCreate(creator, newAccountName, password string, fee *types.Asset) (*OperResp, error) {
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

//AccountCreateDelegation сreate a user account using delegation.
func (client *Client) AccountCreateDelegation(creator, newAccountName, password string, delegated, fee *types.Asset) (*OperResp, error) {
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

	tx := &types.AccountCreateWithDelegationOperation{
		Fee:            fee,
		Delegation:     delegated,
		Creator:        creator,
		NewAccountName: newAccountName,
		Owner:          &owner,
		Active:         &active,
		Posting:        &posting,
		MemoKey:        listKeys["memo"].Public,
		JSONMetadata:   jsonMeta,
		Extensions:     []interface{}{},
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(creator, trx)
	return &OperResp{NameOper: "AccountCreateDelegation", Bresp: resp}, err
}

//Delegation allows you to delegate a number of GESTS to another user.
func (client *Client) Delegation(from, to string, vestingshares *types.Asset) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.DelegateVestingSharesOperation{
		Delegator:     from,
		Delegatee:     to,
		VestingShares: vestingshares,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(from, trx)
	return &OperResp{NameOper: "Delegation", Bresp: resp}, err
}

//SetWithdrawVestingRoute allows you to redirect a certain percentage of GESTS to another user when the SG is lowered.
func (client *Client) SetWithdrawVestingRoute(from, to string, percent uint16, autovest bool) (*OperResp, error) {
	var trx []types.Operation

	tx := &types.SetWithdrawVestingRouteOperation{
		FromAccount: from,
		ToAccount:   to,
		Percent:     percent,
		AutoVest:    autovest,
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(from, trx)
	return &OperResp{NameOper: "SetWithdrawVestingRoute", Bresp: resp}, err
}

func (client *Client) ChainPropertiesUpdate(owner string, accountcreationfee *types.Asset, maxblocksize uint32, sbdinterestrate uint16) (*OperResp, error) {
	var trx []types.Operation

	param, err1 := client.Database.GetWitnessByAccount(owner)
	if err1 != nil {
		return nil, err1
	}

	tx := &types.WitnessUpdateOperation{
		Owner:           owner,
		URL:             param.URL,
		BlockSigningKey: param.SigningKey,
		Props: &types.ChainProperties{
			AccountCreationFee: accountcreationfee,
			MaximumBlockSize:   maxblocksize,
			SBDInterestRate:    sbdinterestrate,
		},
		Fee: SetAsset(0.000, "STEEM"),
	}

	trx = append(trx, tx)
	resp, err := client.SendTrx(owner, trx)
	return &OperResp{NameOper: "WitnessUpdate", Bresp: resp}, err
}

//AddKeys adds a key to the account
/*
keytype:
0 - active
1 - posting
*/
func (client *Client) AddKeys(username string, keytype int, keys []string) (*OperResp, error) {
	var trx []types.Operation

	var ka = map[string]int64{}
	var aa = map[string]int64{}
	var weight uint32

	var active *types.Authority
	var posting *types.Authority

	r, e := client.Database.GetAccounts(username)
	if e != nil {
		return nil, e
	}

	memo := r[0].MemoKey
	jsondata := r[0].JSONMetadata

	switch {
	case keytype == 0:
		for k, v := range r[0].Active.AccountAuths {
			aa[k] = v
		}

		for k, v := range r[0].Active.KeyAuths {
			ka[k] = v
		}
		weight = r[0].Active.WeightThreshold
	case keytype == 1:
		for k, v := range r[0].Posting.AccountAuths {
			aa[k] = v
		}

		for k, v := range r[0].Posting.KeyAuths {
			ka[k] = v
		}
		weight = r[0].Posting.WeightThreshold
	default:
		return nil, fmt.Errorf("The types parameter is invalid")
	}

	for _, k := range keys {
		if strings.HasPrefix(k, "GLS") {
			ka[k] = 1
		} else {
			aa[k] = 1
		}
	}

	keylist := &types.Authority{
		WeightThreshold: weight,
		AccountAuths:    aa,
		KeyAuths:        ka,
	}
	switch {
	case keytype == 0:
		active = keylist
	case keytype == 1:
		posting = keylist
	}
	tx := types.AccountUpdateOperation{
		Account:      username,
		Posting:      posting,
		Active:       active,
		MemoKey:      memo,
		JSONMetadata: jsondata,
	}

	trx = append(trx, &tx)

	resp, err := client.SendTrx(username, trx)
	return &OperResp{NameOper: "AddKeys", Bresp: resp}, err
}

//RemoveKeys removes a key from the account
/*
keytype:
0 - active
1 - posting
*/
func (client *Client) RemoveKeys(username string, keytype int, keys []string) (*OperResp, error) {
	var trx []types.Operation

	var ka = map[string]int64{}
	var aa = map[string]int64{}
	var weight uint32

	var active *types.Authority
	var posting *types.Authority

	r, e := client.Database.GetAccounts(username)
	if e != nil {
		return nil, e
	}

	memo := r[0].MemoKey
	jsondata := r[0].JSONMetadata

	switch {
	case keytype == 0:
		if len(r[0].Active.AccountAuths) < 1 && len(r[0].Active.KeyAuths) < 2 {
			return nil, fmt.Errorf("You can not delete a single key")
		}
		weight = r[0].Active.WeightThreshold
		for _, addkey := range keys {
			for k, v := range r[0].Active.AccountAuths {
				if k != addkey {
					aa[k] = v
				}
			}

			for k, v := range r[0].Active.KeyAuths {
				if k != addkey {
					ka[k] = v
				}
			}

		}

	case keytype == 1:
		if len(r[0].Posting.AccountAuths) < 1 && len(r[0].Posting.KeyAuths) < 2 {
			return nil, fmt.Errorf("You can not delete a single key")
		}
		weight = r[0].Posting.WeightThreshold
		for _, addkey := range keys {
			for k, v := range r[0].Posting.AccountAuths {
				if k != addkey {
					aa[k] = v
				}
			}

			for k, v := range r[0].Posting.KeyAuths {
				if k != addkey {
					ka[k] = v
				}
			}

		}

	default:
		return nil, fmt.Errorf("The types parameter is invalid")
	}

	keylist := &types.Authority{
		WeightThreshold: weight,
		AccountAuths:    aa,
		KeyAuths:        ka,
	}
	switch {
	case keytype == 0:
		active = keylist
	case keytype == 1:
		posting = keylist
	}
	tx := types.AccountUpdateOperation{
		Account:      username,
		Posting:      posting,
		Active:       active,
		MemoKey:      memo,
		JSONMetadata: jsondata,
	}

	trx = append(trx, &tx)

	resp, err := client.SendTrx(username, trx)
	return &OperResp{NameOper: "RemoveKeys", Bresp: resp}, err
}

package client

import (
	"time"

	"github.com/asuleymanov/steem-go/api/network_broadcast"
	"github.com/asuleymanov/steem-go/transactions"
	"github.com/asuleymanov/steem-go/types"
)

//SendTrx generates and sends an array of transactions to GOLOS.
func (client *Client) SendTrx(username string, strx []types.Operation) (*BResp, error) {
	var bresp BResp

	// Getting the necessary parameters
	props, err := client.Database.GetDynamicGlobalProperties()
	if err != nil {
		return nil, err
	}

	// Creating a Transaction
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return nil, err
	}
	tx := transactions.NewSignedTransaction(&types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	})

	// Adding Operations to a Transaction
	for _, val := range strx {
		tx.PushOperation(val)
	}

	expTime := time.Now().Add(59 * time.Minute).UTC()
	tm := types.Time{
		Time: &expTime,
	}
	tx.Expiration = &tm

	// Obtain the key required for signing
	privKeys, err := client.SigningKeys(strx[0])
	if err != nil {
		return nil, err
	}

	// Sign the transaction
	if err := tx.Sign(privKeys, client.chainID); err != nil {
		return nil, err
	}

	// Sending a transaction
	var resp *network_broadcast.BroadcastResponse
	if client.AsyncProtocol {
		err = client.NetworkBroadcast.BroadcastTransaction(tx.Transaction)
	} else {
		resp, err = client.NetworkBroadcast.BroadcastTransactionSynchronous(tx.Transaction)
	}

	bresp.JSONTrx, _ = JSONTrxString(tx)

	if err != nil {
		return &bresp, err
	}
	if resp != nil && !client.AsyncProtocol {
		bresp.ID = resp.ID
		bresp.BlockNum = resp.BlockNum
		bresp.TrxNum = resp.TrxNum
		bresp.Expired = resp.Expired

		return &bresp, nil
	}

	return &bresp, nil
}

func (client *Client) GetTrx(strx []types.Operation) (*types.Transaction, error) {
	// Getting the necessary parameters
	props, err := client.Database.GetDynamicGlobalProperties()
	if err != nil {
		return nil, err
	}

	// Creating a Transaction
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return nil, err
	}
	tx := &types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	}

	// Adding Operations to a Transaction
	for _, val := range strx {
		tx.PushOperation(val)
	}

	expTime := time.Now().Add(59 * time.Minute).UTC()
	tm := types.Time{
		Time: &expTime,
	}
	tx.Expiration = &tm

	return tx, nil
}

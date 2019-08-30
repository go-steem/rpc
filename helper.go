package client

import (
	"strconv"

	"github.com/asuleymanov/steem-go/transactions"
	"github.com/asuleymanov/steem-go/types"
)

//SetAsset returns data of type Asset
func SetAsset(amount float64, symbol string) *types.Asset {
	return &types.Asset{Amount: amount, Symbol: symbol}
}

//PerMvest returns the ratio of TotalVersingFund to TotalVestingShares.
func (client *Client) PerMvest() (float64, error) {
	dgp, errdgp := client.API.GetDynamicGlobalProperties()
	if errdgp != nil {
		return 0, errdgp
	}

	tvfs := dgp.TotalVersingFund.Amount
	tvs := dgp.TotalVestingShares.Amount

	spmtmp := (tvfs / tvs) * 1000000

	str := strconv.FormatFloat(spmtmp, 'f', 3, 64)
	spm, errspm := strconv.ParseFloat(str, 64)
	if errspm != nil {
		return 0, errspm
	}

	return spm, nil
}

//JSONTrxString generate Trx to String
func JSONTrxString(v *transactions.SignedTransaction) (string, error) {
	ans, err := types.JSONMarshal(v)
	if err != nil {
		return "", err
	}
	return string(ans), nil
}

//JSONOpString generate Operations to String
func JSONOpString(v []types.Operation) (string, error) {
	var tx types.Operations

	tx = append(tx, v...)

	ans, err := types.JSONMarshal(tx)
	if err != nil {
		return "", err
	}
	return string(ans), nil
}

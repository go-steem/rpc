package client

import (
	"strconv"
	"strings"
)

//SbdMedianPrice returns the average cost of GBG when converting GOLOS.
func (client *Client) SbdMedianPrice() (float64, error) {
	smpreq, errsmp := client.Database.GetFeedHistory()
	if errsmp != nil {
		return 0, errsmp
	}

	base, errbase := strconv.ParseFloat(strings.Split(smpreq.CurrentMedianHistory.Base, " ")[0], 64)
	if errbase != nil {
		return 0, errbase
	}

	quote, errquote := strconv.ParseFloat(strings.Split(smpreq.CurrentMedianHistory.Quote, " ")[0], 64)
	if errquote != nil {
		return 0, errquote
	}

	smptmp := base / quote
	str := strconv.FormatFloat(smptmp, 'f', 3, 64)

	smp, errsmp := strconv.ParseFloat(str, 64)
	if errsmp != nil {
		return 0, errsmp
	}

	return smp, nil
}

//SteemPerMvest returns the ratio of TotalVersingFundSteem to TotalVestingShares.
func (client *Client) SteemPerMvest() (float64, error) {
	dgp, errdgp := client.Database.GetDynamicGlobalProperties()
	if errdgp != nil {
		return 0, errdgp
	}

	tvfs, errtvfs := strconv.ParseFloat(strings.Split(dgp.TotalVersingFundSteem, " ")[0], 64)
	if errtvfs != nil {
		return 0, errtvfs
	}

	tvs, errtvs := strconv.ParseFloat(strings.Split(dgp.TotalVestingShares, " ")[0], 64)
	if errtvs != nil {
		return 0, errtvs
	}

	spmtmp := (tvfs / tvs) * 1000000
	str := strconv.FormatFloat(spmtmp, 'f', 3, 64)

	spm, errspm := strconv.ParseFloat(str, 64)
	if errspm != nil {
		return 0, errspm
	}

	return spm, nil
}

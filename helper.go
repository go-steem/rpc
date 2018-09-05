package client

import (
	"errors"
	"time"

	"github.com/asuleymanov/steem-go/types"
)

//FollowersList returns the subscriber's list of subscribers
func (client *Client) FollowersList(username string) ([]string, error) {
	var followers []string
	fc, _ := client.Follow.GetFollowCount(username)
	fccount := fc.FollowerCount
	i := 0
	for i < fccount {
		req, err := client.Follow.GetFollowers(username, "", "blog", 1000)
		if err != nil {
			return followers, err
		}

		for _, v := range req {
			followers = append(followers, v.Follower)
		}
		i = i + 1000
	}

	return followers, nil
}

//FollowingList returns the list of user subscriptions
func (client *Client) FollowingList(username string) ([]string, error) {
	var following []string
	fc, _ := client.Follow.GetFollowCount(username)
	fccount := fc.FollowingCount
	i := 0
	for i < fccount {
		req, err := client.Follow.GetFollowing(username, "", "blog", 100)
		if err != nil {
			return following, err
		}

		for _, v := range req {
			following = append(following, v.Following)
		}
		i = i + 100
	}

	return following, nil
}

//GetVotingPower returns the POWER of the user based on the time of the last vote.
func (client *Client) GetVotingPower(username string) (int, error) {
	conf, errc := client.Database.GetConfig()
	if errc != nil {
		return 0, errc
	}

	acc, erra := client.Database.GetAccounts([]string{username})
	if erra != nil {
		return 0, erra
	}

	vp := acc[0].VotingPower
	lvt := acc[0].LastVoteTime
	dtn := time.Now()

	regen := conf.Steemit100Percent * int(dtn.Sub(*lvt.Time).Seconds()) / conf.SteemitVoteRegenerationSeconds
	power := (vp + regen) // 100
	if power > 10000 {
		power = 10000
	}
	return power, nil
}

//GetAuthorReward returns information about the reward received for publication.
func (client *Client) GetAuthorReward(username, permlink string, full bool) (*types.AuthorRewardOperation, error) {
	if !full {
		resp, err := client.Database.GetAccountHistory(username, -1, 1000)
		if err != nil {
			return nil, err
		}
		for k, v := range resp {
			if v.OperationType == "author_reward" {
				op := resp[k].Operation.Data()
				if op.(*types.AuthorRewardOperation).Permlink == permlink {
					return op.(*types.AuthorRewardOperation), nil
				}
			}
		}
		return nil, errors.New("In the last 1000 entries from the user's history, no data was found")
	}
	from := 1000
	limit := 1000
	var lastBlock uint32
	for {
		ans, err := client.Database.GetAccountHistory(username, int64(from), uint32(limit))
		if err != nil {
			return nil, err
		}

		if len(ans) == 0 {
			break
		}
		for k, v := range ans {
			if v.OperationType == "author_reward" {
				op := ans[k].Operation.Data()
				if op.(*types.AuthorRewardOperation).Permlink == permlink {
					return op.(*types.AuthorRewardOperation), nil
				}
			}
		}
		s := ans[len(ans)-1:]
		block := s[0].BlockNumber
		if block == lastBlock {
			break
		}

		lastBlock = block
		from = from + limit
	}
	return nil, errors.New("Data about the publication is not found in the entire history of the user's actions")
}

//GetCommentOptionsOperation generates CommentOptionsOperation depending on the incoming data
func GetCommentOptionsOperation(username, permlink string, options PCOptions) *types.CommentOptionsOperation {
	var ext []interface{}
	var AV, ACR bool
	symbol := "GBG"
	MAP := "1000000.000 " + symbol
	PSD := options.Percent
	Extens := []interface{}{}

	if options.Percent == 0 {
		MAP = "0.000 " + symbol
		PSD = 10000
	} else if options.Percent == 50 {
		PSD = 10000
	} else {
		PSD = 0
	}

	if options.AllowVotes == nil || *options.AllowVotes {
		AV = OptionsTrue
	}

	if options.AllowCurationRewards == nil || *options.AllowCurationRewards {
		ACR = OptionsTrue
	}

	if options.BeneficiarieList != nil && len(*options.BeneficiarieList) > 0 {
		var benList []types.Beneficiarie
		var benef types.CommentPayoutBeneficiaries
		for _, val := range *options.BeneficiarieList {
			benList = append(benList, types.Beneficiarie{val.Account, val.Weight})
		}
		benef.Beneficiaries = benList
		ext = append(ext, 0)
		ext = append(ext, benef)
	}

	if len(ext) > 0 {
		Extens = []interface{}{ext}
	}

	return &types.CommentOptionsOperation{
		Author:               username,
		Permlink:             permlink,
		MaxAcceptedPayout:    MAP,
		PercentSteemDollars:  PSD,
		AllowVotes:           AV,
		AllowCurationRewards: ACR,
		Extensions:           Extens,
	}
}

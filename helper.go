package rpc

import (
	"errors"
	"time"

	"github.com/asuleymanov/rpc/types"
)

func (api *Client) Followers_List(username string) ([]string, error) {
	var followers []string
	fc, errfc := api.Follow.GetFollowCount(username)
	if errfc != nil {
		return followers, errfc
	} else {
		fccount := fc.FollowerCount
		i := 0
		for i < fccount {
			req, err := api.Follow.GetFollowers(username, "", "blog", 1000)
			if err != nil {
				return followers, err
			}

			for _, v := range req {
				followers = append(followers, v.Follower)
			}
			i = i + 1000
		}
	}

	return followers, nil
}

func (api *Client) Following_List(username string) ([]string, error) {
	var following []string
	fc, errfc := api.Follow.GetFollowCount(username)
	if errfc != nil {
		return following, errfc
	} else {
		fccount := fc.FollowingCount
		i := 0
		for i < fccount {
			req, err := api.Follow.GetFollowing(username, "", "blog", 100)
			if err != nil {
				return following, err
			}

			for _, v := range req {
				following = append(following, v.Following)
			}
			i = i + 100
		}
	}

	return following, nil
}

func (api *Client) GetVotingPower(username string) (int, error) {
	conf, errc := api.Database.GetConfig()
	if errc != nil {
		return 0, errc
	}

	acc, erra := api.Database.GetAccounts([]string{username})
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

func (api *Client) GetAuthorReward(username, permlink string, full bool) (*types.AuthorRewardOperation, error) {
	if !full {
		resp, err := api.Database.GetAccountHistory(username, -1, 1000)
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
		return nil, errors.New("In the last 1000 entries from the user's history, no data was found.")
	} else {
		from := 1000
		limit := 1000
		var lastBlock uint32 = 0
		for {
			ans, err := api.Database.GetAccountHistory(username, int64(from), uint32(limit))
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
		return nil, errors.New("Data about the publication is not found in the entire history of the user's actions.")
	}
}

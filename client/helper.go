package client

func (api *Client) Followers_List(username string) ([]string, error) {
	var followers []string
	fc, errfc := api.Rpc.Follow.GetFollowCount(username)
	if errfc != nil {
		return followers, errfc
	} else {
		fccount := fc.FollowerCount
		i := 0
		for i < fccount {
			req, err := api.Rpc.Follow.GetFollowers(username, "", "blog", 1000)
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
	fc, errfc := api.Rpc.Follow.GetFollowCount(username)
	if errfc != nil {
		return following, errfc
	} else {
		fccount := fc.FollowingCount
		i := 0
		for i < fccount {
			req, err := api.Rpc.Follow.GetFollowing(username, "", "blog", 100)
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

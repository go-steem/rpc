package client

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
)

//VerifyVoterWeight check whether there is a voter on the list of those who have already voted for the weight of the vote.
func (client *Client) VerifyVoterWeight(author, permlink, voter string, weight int) (bool, error) {
	ans, err := client.Database.GetActiveVotes(author, permlink)

	if err != nil {
		return false, errors.Wrapf(err, "Error Verify Voter: ")
	}
	for _, v := range ans {
		if v.Voter == voter && v.Percent == weight {
			return true, nil
		}
	}
	return false, nil
}

//VerifyVoter check whether there is a voter on the list of those who have already voted without taking into account the weight of the vote.
func (client *Client) VerifyVoter(author, permlink, voter string) (bool, error) {
	ans, err := client.Database.GetActiveVotes(author, permlink)

	if err != nil {
		return false, errors.Wrapf(err, "Error Verify Voter: ")
	}
	for _, v := range ans {
		if v.Voter == voter {
			return true, nil
		}
	}
	return false, nil
}

//VerifyVotes check whether there are voted
func (client *Client) VerifyVotes(author, permlink string) (bool, error) {
	ans, err := client.Database.GetActiveVotes(author, permlink)

	if err != nil {
		return false, errors.Wrapf(err, "Error Verify Votes: ")
	}
	if len(ans) > 0 {
		return true, nil
	}
	return false, nil
}

//ExistComments check whether the entry in GOLOS is a comment.
func (client *Client) ExistComments(author, permlink string) (bool, error) {
	ans, err := client.Database.GetContentReplies(author, permlink)

	if err != nil {
		return false, errors.Wrapf(err, "Error Exist Comments: ")
	}
	if len(ans) > 0 {
		return true, nil
	}
	return false, nil
}

//VerifyReblogs сheck if the user made a repost entry in GOLOS
func (client *Client) VerifyReblogs(author, permlink, rebloger string) (bool, error) {
	ans, err := client.Follow.GetRebloggedBy(author, permlink)
	if err != nil {
		return false, errors.Wrapf(err, "Error Verify Reblogs: ")
	}
	for _, v := range ans {
		if *v == rebloger {
			return true, nil
		}
	}
	return false, nil
}

//VerifyFollow сheck if one user is signed for the second in GOLOS
func (client *Client) VerifyFollow(follower, following string) (bool, error) {
	ans, err := client.Follow.GetFollowing(follower, following, "blog", 1)
	if err != nil {
		return false, errors.Wrapf(err, "Error Verify Follow: ")
	}
	for _, v := range ans {
		if (v.Follower == follower) && (v.Following == following) {
			return true, nil
		}
	}
	return false, nil
}

//VerifyPost сheck if there is an entry in GOLOS
func (client *Client) VerifyPost(author, permlink string) (bool, error) {
	ans, err := client.Database.GetContent(author, permlink)

	if err != nil {
		return false, errors.Wrapf(err, "Error Verify Post: ")
	}
	if (ans.Author == author) && (ans.Permlink == permlink) {
		return true, nil
	}
	return false, nil
}

//VerifyDelegatePostingKeySign check whether the user has delegated the opportunity to use his post by using operations from a given user.
func (client *Client) VerifyDelegatePostingKeySign(fromUser, toUser string) (bool, error) {
	acc, err := client.Database.GetAccounts([]string{fromUser})
	if err != nil {
		return false, errors.Wrapf(err, "Error Verify Delegate Vote Sign: ")
	}

	if len(acc) == 1 {
		for _, v := range acc[0].Posting.AccountAuths {
			tu := strings.Split(strings.Replace(strings.Replace(fmt.Sprintf("%v", v), "[", "", -1), "]", "", -1), " ")
			if tu[0] == toUser {
				if tu[1] == fmt.Sprintf("%v", acc[0].Posting.WeightThreshold) {
					return true, nil
				}
			}
		}
	}
	return false, nil
}

//VerifyFirstPost сheck whether the post of the user is his first post in GOLOS
func (client *Client) VerifyFirstPost(username string) (bool, error) {
	d := time.Now()
	cont, err := client.Database.GetDiscussionsByAuthorBeforeDate(username, "", d.Format("2006-01-02T00:00:00"), 100)
	if err != nil {
		return false, errors.Wrapf(err, "Error Verify First Post: ")
	}

	return len(cont) <= 1, nil
}

//VerifyUser сheck if the user exists in GOLOS
func (client *Client) VerifyUser(username string) (bool, error) {
	acc, err := client.Database.GetAccounts([]string{username})
	if err != nil {
		return false, errors.Wrapf(err, "Error Verify User: ")
	}

	return len(acc) == 1, nil
}

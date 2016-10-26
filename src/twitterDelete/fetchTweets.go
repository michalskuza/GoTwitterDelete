package twitterDelete

import "github.com/dghubble/go-twitter/twitter"

// GetAllTweetsList --
func GetAllTweetsList(client twitter.Client) []twitter.Tweet {
	var tweets []twitter.Tweet
	tmp := GetTweetsLessThanMaxID(client, nil)
	tweets = append(tweets, tmp...)
	maxID := tweets[len(tweets)-1].ID

	var lastMaxID int64

	for maxID != lastMaxID {
		tmp = GetTweetsLessThanMaxID(client, &maxID)
		tweets = append(tweets, tmp[1:]...)
		lastMaxID = maxID
		maxID = tmp[len(tmp)-1].ID
	}

	return tweets
}

//GetTweetsLessThanMaxID --
func GetTweetsLessThanMaxID(client twitter.Client, maxID *int64) []twitter.Tweet {
	var userTimelineParams *twitter.UserTimelineParams

	if maxID != nil {
		userTimelineParams = &twitter.UserTimelineParams{MaxID: *maxID, Count: 20}
	} else {
		userTimelineParams = &twitter.UserTimelineParams{Count: 20}
	}

	tweets, _, _ := client.Timelines.UserTimeline(userTimelineParams)

	return tweets
}

package twitterDelete

import "github.com/dghubble/go-twitter/twitter"

// DeleteAllTweetsExceptFavorited --
func DeleteAllTweetsExceptFavorited(client twitter.Client, tweets []twitter.Tweet) {
	for _, tweet := range tweets {
		if !tweet.Favorited {
			DeleteTweet(client, tweet.ID)
		}
	}
}

// DeleteAllTweets --
func DeleteAllTweets(client twitter.Client, tweets []twitter.Tweet) {
	for _, tweet := range tweets {
		DeleteTweet(client, tweet.ID)
	}
}

// DeleteTweet --
func DeleteTweet(client twitter.Client, tweedID int64) {
	client.Statuses.Destroy(tweedID, nil)
}

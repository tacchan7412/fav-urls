package main

import (
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

var api *anaconda.TwitterApi
var twitterHost = "twitter.com"

func newTwitterAPI() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
}

func extractURLsFromTweets(tweets []anaconda.Tweet) (urls []string) {
	for _, tweet := range tweets {
		urls = append(urls, extractURLsFromTweet(tweet)...)
	}
	return
}

func extractURLsFromTweet(tweet anaconda.Tweet) (urls []string) {
	entitiesUrls := tweet.Entities.Urls
	for _, entitiesUrl := range entitiesUrls {
		u, err := url.Parse(entitiesUrl.Expanded_url)
		if err != nil {
			continue
		}
		if u.Host == twitterHost {
			continue
		}
		urls = append(urls, u.String())
	}
	return
}

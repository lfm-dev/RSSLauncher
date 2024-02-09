package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mmcdole/gofeed"
)

func getFeeds(feedsUrls []string) []Feed {
	feeds := make([]Feed, 0)

	feedParser := gofeed.NewParser()
	for i, feedUrl := range feedsUrls {
		goFeed, err := feedParser.ParseURL(feedUrl)
		if err != nil {
			fmt.Printf("Can't get %s data\n", feedUrl)
			continue
		}

		feedItems := getFeedItems(goFeed)
		feed := Feed{
			url:        feedsUrls[i],
			name:       goFeed.Title,
			lastUpdate: feedItems[0].date, // date of newest post
			items:      feedItems,
		}
		feeds = append(feeds, feed)
	}

	return feeds
}

func getFeedsUrl() []string {
	homePath, _ := os.UserHomeDir()
	feedsFilePath := homePath + "/.config/RSS/feeds.txt"
	feeds, err := os.ReadFile(feedsFilePath)
	if err != nil {
		fmt.Println("oops")
	}
	feedsUrls := strings.Split(strings.TrimSpace(string(feeds)), "\n")
	return feedsUrls
}

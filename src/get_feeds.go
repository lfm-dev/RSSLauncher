package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/schollz/progressbar/v3"
)

func getFeedsUrl() []string {
	feeds, err := os.ReadFile(feedsFilePath)
	if err != nil {
		panic(err)
	}
	feedsUrls := strings.Split(strings.TrimSpace(string(feeds)), "\n")
	return feedsUrls
}

//TODO can you update feeds with goroutines?
func getFeeds(feedsUrls []string) []Feed {
	fmt.Printf("Updating %d feeds...\n", len(feedsUrls))
	progressBar := progressbar.Default(int64(len(feedsUrls)))
	feeds := make([]Feed, 0)

	feedParser := gofeed.NewParser()
	for i, feedUrl := range feedsUrls {
		goFeed, err := feedParser.ParseURL(feedUrl)
		if err != nil {
			fmt.Printf("\nError: Can't get %s data\n", feedUrl)
			time.Sleep(2000 * time.Millisecond)
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
		progressBar.Add(1)
	}

	return feeds
}

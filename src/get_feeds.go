package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/schollz/progressbar/v3"
)

func getFeedItems(goFeed *gofeed.Feed) []FeedItem {
	feedItems := make([]FeedItem, 0)
	for _, item := range goFeed.Items {
		if postHasIgnoredWord(item.Title) {
			continue
		}
		feedItem := FeedItem{
			url:          item.Link,
			title:        item.Title,
			date:         *item.PublishedParsed,
			dateFormated: item.PublishedParsed.Format("02-01-2006"),
		}
		feedItems = append(feedItems, feedItem)
	}

	return feedItems
}

func getFeedsUrls() []string {
	feeds, err := os.ReadFile(feedsFilePath)
	if err != nil {
		panic(errors.New("Can't read feeds.txt"))
	}
	feedsUrls := strings.Split(strings.TrimSpace(string(feeds)), "\n")
	return feedsUrls
}

//TODO can you update feeds with goroutines?
func getFeeds() []Feed {
	feedsUrls := getFeedsUrls()
	fmt.Printf("Updating %d feeds...\n", len(feedsUrls))
	progressBar := progressbar.Default(int64(len(feedsUrls)))
	feeds := make([]Feed, 0)

	feedParser := gofeed.NewParser()
	for _, feedUrl := range feedsUrls {
		goFeed, err := feedParser.ParseURL(feedUrl)
		if err != nil {
			fmt.Printf("\nError: Can't get %s data\n", feedUrl)
			time.Sleep(2000 * time.Millisecond)
			continue
		}

		feedItems := getFeedItems(goFeed)
		feed := Feed{
			feedUrl:    goFeed.FeedLink,
			url:        goFeed.Link,
			name:       goFeed.Title,
			lastUpdate: feedItems[0].date, // date of newest post
			items:      feedItems,
		}
		feeds = append(feeds, feed)
		progressBar.Add(1)
	}

	return feeds
}

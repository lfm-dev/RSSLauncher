package main

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/schollz/progressbar/v3"
)

func getFeedItems(goFeed *gofeed.Feed) []FeedItem {
	feedItems := make([]FeedItem, 0)
	for _, item := range goFeed.Items {
		if itemHasIgnoredWord(item.Title) {
			continue
		}
		feedItem := FeedItem{
			url:          item.Link,
			title:        item.Title,
			dateFormated: item.PublishedParsed.Format("02-01-2006"),
			read:         false, // for now
		}
		feedItems = append(feedItems, feedItem)
	}
	return feedItems
}

//TODO can you update feeds with goroutines?
func getFeeds() []Feed {
	feedsUrls := getFileLines(feedsFilePath)
	fmt.Printf("Updating %d feeds...\n", len(feedsUrls))
	progressBar := progressbar.Default(int64(len(feedsUrls)))
	feeds := make([]Feed, 0)

	feedParser := gofeed.NewParser()
	for _, feedUrl := range feedsUrls {
		goFeed, err := feedParser.ParseURL(feedUrl)
		if err != nil {
			fmt.Printf("\nError: Can't get %s data\n", feedUrl)
			time.Sleep(2000 * time.Millisecond) // so the user can read it
			progressBar.Add(1)
			continue
		}

		feedItems := getFeedItems(goFeed)
		feed := Feed{
			feedUrl: goFeed.FeedLink,
			url:     goFeed.Link,
			name:    goFeed.Title,
			items:   feedItems,
		}
		feeds = append(feeds, feed)
		progressBar.Add(1)
	}
	return feeds
}

package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/schollz/progressbar/v3"
)

func getNewFeedItems(goFeed *gofeed.Feed, feedUrl string) []FeedItem {
	feedItems := make([]FeedItem, 0)
	for _, item := range goFeed.Items {
		if itemHasIgnoredWord(item.Title) {
			continue
		}
		feedItem := FeedItem{
			FeedUrl: feedUrl,
			ItemUrl: item.Link,
			Title:   item.Title,
			Date:    *item.PublishedParsed,
			Read:    false,
		}
		feedItems = append(feedItems, feedItem)
	}
	return feedItems
}

//TODO can you update feeds with goroutines?
func getFeeds() []Feed {
	feedsUrls := getFileLines(feedsFilePath)

	fmt.Printf("Updating %d feeds...\n", len(feedsUrls)) // len(feedsUrls) includes categories
	progressBar := progressbar.Default(int64(len(feedsUrls)))
	feeds := make([]Feed, 0)

	feedCategory := "noCategory"
	feedParser := gofeed.NewParser()
	for _, feedUrl := range feedsUrls {

		if strings.HasPrefix(feedUrl, "#") {
			feedCategory = feedUrl[1:]
			continue
		}

		goFeed, err := feedParser.ParseURL(feedUrl)
		if err != nil {
			fmt.Printf("\nError: Can't get %s data\n", feedUrl)
			time.Sleep(1000 * time.Millisecond) // so the user can read it
			progressBar.Add(1)
			continue
		}

		newFeedItems := getNewFeedItems(goFeed, feedUrl)
		addNewItemsToDB(newFeedItems)
		allFeedItems := getAllItemsFromDB(feedUrl)
		sortItemsByDate(allFeedItems)

		feed := Feed{
			url:      goFeed.Link,
			name:     goFeed.Title,
			category: feedCategory,
			items:    allFeedItems,
		}
		feeds = append(feeds, feed)
		progressBar.Add(1)
	}
	return feeds
}

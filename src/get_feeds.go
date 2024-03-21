package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/schollz/progressbar/v3"
)

func getItemsFromFeed(goFeed *gofeed.Feed, feedUrl string) []FeedItem {
	feedItems := make([]FeedItem, 0)

	for _, item := range goFeed.Items {

		if itemHasBlacklistedWord(item.Title) {
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

func getFeedsNumber(feedFileLines []string) int {
	nFeeds := 0
	for _, line := range feedFileLines {
		if !strings.HasPrefix(line, "#") { // # category name
			nFeeds++
		}
	}
	return nFeeds
}

//TODO can you update feeds with goroutines?
func getFeeds() []Feed {
	feedsFileLines := getFileLines(feedsFilePath)
	nFeeds := getFeedsNumber(feedsFileLines)

	fmt.Printf("Updating %d feeds...\n", nFeeds)
	progressBar := progressbar.Default(int64(nFeeds))
	feeds := make([]Feed, 0)

	feedCategory := "noCategory"
	feedParser := gofeed.NewParser()
	for _, line := range feedsFileLines {

		if strings.HasPrefix(line, "#") {
			feedCategory = line[1:]
			continue
		}

		feedUrl := line
		goFeed, err := feedParser.ParseURL(feedUrl)
		if err != nil {
			fmt.Printf("\nError: Can't get %s data\n", feedUrl)
			time.Sleep(1000 * time.Millisecond) // so the user can read it
			progressBar.Add(1)
			continue
		}

		newFeedItems := getItemsFromFeed(goFeed, feedUrl)
		addNewItemsToDB(newFeedItems)
		allFeedItems := getAllItemsFromDB(feedUrl)

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

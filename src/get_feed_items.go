package main

import "github.com/mmcdole/gofeed"

func getFeedItems(goFeed *gofeed.Feed) []FeedItem {
	feedItems := make([]FeedItem, 0)
	for _, item := range goFeed.Items {
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

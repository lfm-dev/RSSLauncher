package main

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

func getFeed(feedUrl string, fp *gofeed.Parser) *gofeed.Feed {
	feed, err := fp.ParseURL(feedUrl)
	if err != nil {
		fmt.Printf("Can't get %s data\n", feedUrl)
	}
	return feed
}

func main() {
	feedsUrls := make([]string, 0)
	feedsUrls = append(feedsUrls, "https://www.raptitude.com/feed")
	feedsUrls = append(feedsUrls, "https://calnewport.com/blog/feed")
	feedsUrls = append(feedsUrls, "https://go.dev/blog/feed.atom")

	fp := gofeed.NewParser()
	for feedIndex, feedUrl := range feedsUrls {
		feed := getFeed(feedUrl, fp)
		lastUpdate := feed.Items[0].PublishedParsed // update date of newest post
		fmt.Println()
		fmt.Println(feedIndex, feed.Title, lastUpdate)
		for index, item := range feed.Items {
			fmt.Println(index, item.Title, item.PublishedParsed.Format(time.UnixDate))
		}
	}
}

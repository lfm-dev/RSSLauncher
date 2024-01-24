package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

//TODO run command with item url

func getFeed(feedUrl string, fp *gofeed.Parser) (*gofeed.Feed, error) {
	feed, err := fp.ParseURL(feedUrl)
	return feed, err
}

func main() {
	feedsUrls := make([]string, 0)
	feedsUrls = append(feedsUrls, "https://www.raptitude.com/feed")
	feedsUrls = append(feedsUrls, "https://calnewport.com/blog/feed")
	feedsUrls = append(feedsUrls, "https://go.dev/blog/feed.atom")

	fp := gofeed.NewParser()
	for feedIndex, feedUrl := range feedsUrls {
		feed, err := getFeed(feedUrl, fp)
		if err != nil {
			fmt.Printf("Can't get %s data\n", feedUrl)
			continue
		}

		lastUpdate := feed.Items[0].PublishedParsed // update date of newest post
		fmt.Println()
		fmt.Println(feedIndex, feed.Title, lastUpdate)
		for _, item := range feed.Items {
			fmt.Println(item.Title, item.PublishedParsed.Format("02-01-2006"))
		}
	}
}

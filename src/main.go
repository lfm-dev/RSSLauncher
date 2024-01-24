package main

import (
	"fmt"
	"strings"

	"github.com/mmcdole/gofeed"
	"golang.org/x/net/html"
)

//TODO run command with item url

func getFeed(feedUrl string, fp *gofeed.Parser) (*gofeed.Feed, error) {
	feed, err := fp.ParseURL(feedUrl)
	return feed, err
}

func getFeedData(feeds []*gofeed.Feed) {
	for feedIndex, feed := range feeds {
		lastUpdate := feed.Items[0].PublishedParsed // update date of newest post
		fmt.Println()
		fmt.Println(feedIndex, feed.Title, lastUpdate)
		for _, item := range feed.Items {
			fmt.Println(item.Title, item.PublishedParsed.Format("02-01-2006"))
			fmt.Println(item.Content)

			doc, err := html.Parse(strings.NewReader(item.Content))
			if err != nil {
				fmt.Println("Error parsing HTML:", err)
				return
			}

			fmt.Printf("Parsed HTML document: %+v\n", doc)

		}
	}
}

// func renderNode(n *html.Node) string {
// 	var buf bytes.Buffer
// 	w := io.Writer(&buf)
// 	html.Render(w, n)
// 	return buf.String()
// }

func main() {
	feedsUrls := make([]string, 0)
	feedsUrls = append(feedsUrls, "https://www.raptitude.com/feed")
	feedsUrls = append(feedsUrls, "https://calnewport.com/blog/feed")
	feedsUrls = append(feedsUrls, "https://go.dev/blog/feed.atom")

	feeds := make([]*gofeed.Feed, 0)
	fp := gofeed.NewParser()
	for _, feedUrl := range feedsUrls {
		feed, err := getFeed(feedUrl, fp)
		if err != nil {
			fmt.Printf("Can't get %s data\n", feedUrl)
			continue
		}
		feeds = append(feeds, feed)
	}
	getFeedData(feeds)
}

package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/mmcdole/gofeed"
	"github.com/rivo/tview"
)

//TODO run command with item url
//TODO open item on browser on enter

type Feed struct {
	url        string
	name       string
	lastUpdate time.Time
	items      []FeedItem
}

type FeedItem struct {
	url          string
	title        string
	date         time.Time
	dateFormated string
}

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

func main() {
	feedsUrls := make([]string, 0)
	feedsUrls = append(feedsUrls, "https://www.raptitude.com/feed")
	feedsUrls = append(feedsUrls, "https://go.dev/blog/feed.atom")
	feedsUrls = append(feedsUrls, "https://calnewport.com/blog/feed")

	feeds := getFeeds(feedsUrls)
	fmt.Println(feeds)

	//tview
	app := tview.NewApplication()
	table := tview.NewTable().
		SetSelectable(true, false)

	_, rows := 2, len(feedsUrls)
	for row := 0; row < rows; row++ {
		table.SetCell(row, 0,
			tview.NewTableCell(feedsUrls[row]).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(row, 1,
			tview.NewTableCell("").
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter))
	}

	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
	}).SetSelectedFunc(func(row int, column int) {
		fmt.Println(table.GetCell(row, column).Text)
		fmt.Println("ea", feedsUrls[row])
	})

	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}

}

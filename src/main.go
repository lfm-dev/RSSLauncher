package main

import (
	"fmt"
	"os"
	"strings"
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

func getFeedsUrl() []string {
	homePath, _ := os.UserHomeDir()
	feedsFilePath := homePath + "/.config/RSS/feeds.txt"
	feeds, err := os.ReadFile(feedsFilePath)
	if err != nil {
		fmt.Println("oops")
	}
	feedsUrls := strings.Split(strings.TrimSpace(string(feeds)), "\n")
	return feedsUrls
}

func view(feeds []Feed) {

	app := tview.NewApplication()
	table := tview.NewTable().
		SetSelectable(true, false)

	_, rows := 2, len(feeds)
	for row := 0; row < rows; row++ {
		fmt.Println(row)
		table.SetCell(row, 0,
			tview.NewTableCell(feeds[row].name).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		// table.SetCell(row, 1,
		// 	tview.NewTableCell("").
		// 		SetTextColor(tcell.ColorWhite).
		// 		SetAlign(tview.AlignCenter))
	}

	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
	}).SetSelectedFunc(func(row int, column int) {
		fmt.Println(table.GetCell(row, column).Text)
	})

	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}

func main() {
	feedsUrls := getFeedsUrl()
	feeds := getFeeds(feedsUrls)
	view(feeds)
}

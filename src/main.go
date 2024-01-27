package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/mmcdole/gofeed"
	"github.com/rivo/tview"
)

//TODO run command with item url

func getFeed(feedUrl string, fp *gofeed.Parser) (*gofeed.Feed, error) {
	feed, err := fp.ParseURL(feedUrl)
	return feed, err
}

func main() {
	feedsUrls := make([]string, 0)
	feedsUrls = append(feedsUrls, "https://www.raptitude.com/feed")
	feedsUrls = append(feedsUrls, "https://go.dev/blog/feed.atom")
	feedsUrls = append(feedsUrls, "https://calnewport.com/blog/feed")

	// fp := gofeed.NewParser()
	// for feedIndex, feedUrl := range feedsUrls {
	// 	feed, err := getFeed(feedUrl, fp)
	// 	if err != nil {
	// 		fmt.Printf("Can't get %s data\n", feedUrl)
	// 		continue
	// 	}

	// 	lastUpdate := feed.Items[0].PublishedParsed // update date of newest post
	// 	fmt.Println()
	// 	fmt.Println(feedIndex, feed.Title, lastUpdate)
	// 	for _, item := range feed.Items {

	// 		fmt.Println("\n"+item.Title, item.PublishedParsed.Format("02-01-2006")+"\n")

	// 		doc := soup.HTMLParse(item.Content)
	// 		fullText := strings.Split(doc.FullText(), "\n")
	// 		for _, line := range fullText {
	// 			if len(strings.TrimSpace(line)) != 0 {
	// 				fmt.Println(line)
	// 			}
	// 		}
	// 	}
	// }

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
		fmt.Println(feedsUrls[row])
	})

	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}

}

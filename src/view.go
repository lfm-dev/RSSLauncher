package main

import (
	"fmt"
	"os/exec"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	mainFlex  = tview.NewFlex()
	feedsFlex = tview.NewFlex()
	postsFlex = tview.NewFlex()
)

const BROWSER = "firefox"

func getTables(feeds []Feed, app *tview.Application) (*tview.Table, *tview.Table) {
	feedsTable := tview.NewTable().SetSelectable(true, false)
	postsTable := tview.NewTable().SetSelectable(true, false)

	feedsTable.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
	})

	feedsTable.SetSelectedFunc(func(row int, column int) {
		app.SetFocus(postsTable)
	})

	feedsTable.SetSelectionChangedFunc(func(feedIndex int, column int) {
		renderPostsTable(postsTable, feeds[feedIndex])
	})

	postsTable.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.SetFocus(feedsTable)
		}
	})

	postsTable.SetSelectedFunc(func(itemIndex int, _ int) {
		feedIndex, _ := feedsTable.GetSelection() // ignore column
		cmdStruct := exec.Command(BROWSER, feeds[feedIndex].items[itemIndex].url)
		cmdStruct.Output()
	})

	return feedsTable, postsTable
}

func renderFeedsTable(feeds []Feed, feedsTable *tview.Table) {
	for i, feed := range feeds {
		feedsTable.SetCell(i, 0,
			tview.NewTableCell(feed.name).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
	}
}

func renderPostsTable(postsTable *tview.Table, feed Feed) {
	postsTable.Clear()
	for i, post := range feed.items {
		itemLine := fmt.Sprintf("(%s) %s", post.dateFormated, post.title)
		postsTable.SetCell(i, 0,
			tview.NewTableCell(itemLine))
	}
}

func view(feeds []Feed) {

	app := tview.NewApplication()
	feedsTable, postsTable := getTables(feeds, app)

	renderFeedsTable(feeds, feedsTable)

	renderPostsTable(postsTable, feeds[0]) // show first feed posts

	feedsFlex.AddItem(feedsTable, 0, 1, false).SetBorder(true)
	postsFlex.AddItem(postsTable, 0, 1, false).SetBorder(true)
	mainFlex.AddItem(feedsFlex, 0, 1, false).AddItem(postsFlex, 0, 3, false)

	if err := app.SetRoot(mainFlex, true).SetFocus(feedsTable).Run(); err != nil {
		panic(err)
	}

}

package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

//TODO configure postsTable

var (
	mainFlex  = tview.NewFlex()
	feedsFlex = tview.NewFlex()
	postsFlex = tview.NewFlex()
)

func getFeedsTable(feeds []Feed) *tview.Table {
	feedsTable := tview.NewTable().
		SetSelectable(true, false)

	rows := len(feeds)
	for row := 0; row < rows; row++ {
		feedsTable.SetCell(row, 0,
			tview.NewTableCell(feeds[row].name).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
	}
	return feedsTable
}

func makePostsTable(postsTable *tview.Table, feed Feed) {
	postsTable.Clear()
	for i, post := range feed.items {
		postsTable.SetCell(i, 0,
			tview.NewTableCell(post.title))
	}
}

func view(feeds []Feed) {

	app := tview.NewApplication()
	feedsTable := getFeedsTable(feeds)
	postsTable := tview.NewTable().SetSelectable(true, false)

	feedsFlex.AddItem(feedsTable, 0, 1, false).SetBorder(true)
	postsFlex.AddItem(postsTable, 0, 1, false).SetBorder(true)

	mainFlex.AddItem(feedsFlex, 0, 1, false)
	mainFlex.AddItem(postsFlex, 0, 2, false)

	feedsTable.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
	}).SetSelectedFunc(func(row int, column int) {
		app.SetFocus(postsTable)
	})

	feedsTable.SetSelectionChangedFunc(func(feedIndex int, column int) {
		makePostsTable(postsTable, feeds[feedIndex]) // show selected feed posts
	})

	makePostsTable(postsTable, feeds[0]) // show first feed posts

	if err := app.SetRoot(mainFlex, true).SetFocus(feedsTable).Run(); err != nil {
		panic(err)
	}

}

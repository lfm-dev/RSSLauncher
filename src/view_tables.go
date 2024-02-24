package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func renderFeedsTable(feeds []Feed, feedsTable *tview.Table) {
	for i, feed := range feeds {
		feedsTable.SetCell(i, 0,
			tview.NewTableCell(feed.name).SetReference(feed))
	}
}

func renderPostsTable(postsTable *tview.Table, feed Feed) {
	postsTable.Clear()
	for i, post := range feed.items {
		postLine := fmt.Sprintf("(%s) %s", post.dateFormated, post.title)
		postsTable.SetCell(i, 0,
			tview.NewTableCell(postLine))
	}
	postsTable.Select(0, 0)
}

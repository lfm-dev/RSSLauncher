package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func renderFeedsTable(feeds []Feed) {
	for i, feed := range feeds {
		feedsTable.SetCell(i, 0,
			tview.NewTableCell(feed.name).SetReference(feed))
	}
}

func renderItemsTable() {
	itemsTable.Clear()
	feedItems := getFeedData().items
	for i, item := range feedItems {
		itemTitle := fmt.Sprintf("(%s) %s", item.dateFormated, item.title)
		itemsTable.SetCell(i, 0,
			tview.NewTableCell(itemTitle).SetReference(item))
	}
	itemsTable.Select(0, 0)
}

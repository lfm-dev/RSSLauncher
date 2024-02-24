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
		itemCell := tview.NewTableCell(itemTitle).SetReference(item)

		if !item.read {
			itemCell.SetTextColor(unreadTextColor).SetBackgroundColor(unreadBkgColor)
		} else {
			itemCell.SetTextColor(readTextColor).SetBackgroundColor(readBkgColor)
		}

		itemsTable.SetCell(i, 0, itemCell)
	}
	itemsTable.Select(0, 0)
}

func markItemAsRead() {
	feed := getFeedData()
	itemIndex, _ := itemsTable.GetSelection()
	feed.items[itemIndex].read = true
	renderItemsTable() // for now
}

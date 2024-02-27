package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func renderFeedsTable(feeds []Feed) {
	for i, feed := range feeds {
		feedsTable.SetCell(i, 0,
			tview.NewTableCell(feed.name).SetReference(feed))
	}
}

func renderItemsTable(selectFirstItem bool) {
	itemsTable.Clear()
	feedItems := getFeedData().items
	for i, item := range feedItems {
		itemTitle := fmt.Sprintf("(%s) %s", item.dateFormated, item.title)
		itemCell := tview.NewTableCell(itemTitle).SetReference(item)

		if item.read {
			itemCell.SetTextColor(tcell.ColorGrey)
		}

		itemsTable.SetCell(i, 0, itemCell)
	}

	if selectFirstItem {
		itemsTable.Select(0, 0)
	}
}

func markItemAsRead() {
	feed := getFeedData()
	itemIndex, _ := itemsTable.GetSelection()
	if !feed.items[itemIndex].read {
		feed.items[itemIndex].read = true
		renderItemsTable(false)
	}
}

func markAllItemsRead() {
	feed := getFeedData()
	for i := range feed.items {
		if !feed.items[i].read {
			feed.items[i].read = true
		}
	}

	renderItemsTable(false)
}

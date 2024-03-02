package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func renderFeedsTable(feeds []Feed) {
	for i, feed := range feeds {
		feedCell := tview.NewTableCell(feed.name).SetReference(feed)
		if !feedHasUnreadItems(feed) {
			feedCell.SetTextColor(TEXT_COLOR_READ)
		}
		feedsTable.SetCell(i, 0, feedCell)
	}
}

func renderItemsTable(selectFirstItem bool) {
	itemsTable.Clear()
	feedItems := getFeedData().items
	for i, item := range feedItems {
		itemTitle := fmt.Sprintf("(%s) %s",
			item.Date.Format("02-01-2006"),
			item.Title,
		)
		itemCell := tview.NewTableCell(itemTitle).SetReference(item)

		if item.Read {
			itemCell.SetTextColor(TEXT_COLOR_READ)
		}

		itemsTable.SetCell(i, 0, itemCell)
	}

	if selectFirstItem {
		itemsTable.Select(0, 0)
		itemsTable.ScrollToBeginning()
	}
}

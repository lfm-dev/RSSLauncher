package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func renderFeedsTable(feeds []Feed) {
	for i, feed := range feeds {
		feedName := feed.name
		if feedHasUnreadItems(feed) {
			feedName = "*" + feed.name
		}
		feedsTable.SetCell(i, 0,
			tview.NewTableCell(feedName).SetReference(feed))
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
			itemCell.SetTextColor(tcell.ColorGrey)
		}

		itemsTable.SetCell(i, 0, itemCell)
	}

	if selectFirstItem {
		itemsTable.Select(0, 0)
		itemsTable.ScrollToBeginning()
	}
}

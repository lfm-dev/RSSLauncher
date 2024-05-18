package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func renderFeedsTable(feeds []Feed) {
	currentCategory := "noCategory"
	row := 0
	for _, feed := range feeds {

		if feed.category != currentCategory {
			categoryCell := tview.NewTableCell(feed.category).SetSelectable(false).SetBackgroundColor(CATEGORY_BKG_COLOR).SetTextColor(CATEGORY_TEXT_COLOR).SetAlign(1).SetExpansion(1)
			feedsTable.SetCell(row, 0, categoryCell)
			currentCategory = feed.category
			row++
		}

		feedCell := tview.NewTableCell(feed.name).SetReference(feed)

		if !feedHasUnreadItems(feed) {
			feedCell.SetTextColor(TEXT_COLOR_READ)
		}
		feedsTable.SetCell(row, 0, feedCell)
		row++
	}

	if getSelectedCell(feedsTable).GetReference() == nil { // if first cell is a category cell -> select second cell
		feedsTable.Select(1, 0)
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
		itemCell := tview.NewTableCell(itemTitle).SetReference(item).SetExpansion(1)

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

package main

import "github.com/rivo/tview"

func getSelectedCell(table *tview.Table) *tview.TableCell {
	cellIndex, _ := table.GetSelection()

	return table.GetCell(cellIndex, 0)
}

func getFeedData() Feed {
	cellRef := getSelectedCell(feedsTable).GetReference()
	feed := cellRef.(Feed)

	return feed
}

func getItemData() FeedItem {
	cellRef := getSelectedCell(itemsTable).GetReference()
	item := cellRef.(FeedItem)

	return item
}

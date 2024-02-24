package main

func getFeedData() Feed {
	feedIndex, _ := feedsTable.GetSelection()
	cellRef := feedsTable.GetCell(feedIndex, 0).GetReference()
	feed := cellRef.(Feed)

	return feed
}

func getItemData() FeedItem {
	itemIndex, _ := itemsTable.GetSelection()
	cellRef := itemsTable.GetCell(itemIndex, 0).GetReference()
	item := cellRef.(FeedItem)

	return item
}

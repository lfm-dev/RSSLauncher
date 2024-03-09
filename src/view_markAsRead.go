package main

func updateFeedReadStatus(feed Feed) {
	if !feedHasUnreadItems(feed) {
		feedCell := getSelectedCell(feedsTable)
		feedCell.SetTextColor(TEXT_COLOR_READ)
	}
}

func markItemAsRead() {
	feed := getFeedData()
	i, _ := itemsTable.GetSelection()
	if !feed.items[i].Read {
		feed.items[i].Read = true
		markAsReadInDB(feed.items[i].ItemUrl)
		renderItemsTable(false)
		updateFeedReadStatus(feed)
	}
}

func markAllFeedItemsAsRead() {
	feed := getFeedData()
	for i := range feed.items {
		if !feed.items[i].Read {
			feed.items[i].Read = true
			markAsReadInDB(feed.items[i].ItemUrl)
		}
	}
	renderItemsTable(false)
	updateFeedReadStatus(feed)
}

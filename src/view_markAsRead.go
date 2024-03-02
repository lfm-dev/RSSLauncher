package main

func markItemAsRead() {
	feed := getFeedData()
	itemIndex, _ := itemsTable.GetSelection()
	if !feed.items[itemIndex].Read {
		feed.items[itemIndex].Read = true
		markAsReadInDB(feed.items[itemIndex].ItemUrl)
		renderItemsTable(false)
		updateFeedReadStatus(feed)
	}
}

func markAllItemsRead() {
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

package main

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

func markAllItemsAsRead() {
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

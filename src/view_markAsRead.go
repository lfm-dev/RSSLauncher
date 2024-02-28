package main

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

package main

func getFeedData() Feed {
	feedIndex, _ := feedsTable.GetSelection()
	cellRef := feedsTable.GetCell(feedIndex, 0).GetReference()
	feed := cellRef.(Feed)

	return feed
}

func getPostData() FeedItem {
	postIndex, _ := postsTable.GetSelection()
	cellRef := postsTable.GetCell(postIndex, 0).GetReference()
	post := cellRef.(FeedItem)

	return post
}

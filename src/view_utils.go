package main

func getFeedData() Feed {
	feedIndex, _ := feedsTable.GetSelection()
	cellRef := feedsTable.GetCell(feedIndex, 0).GetReference()
	feed := cellRef.(Feed)

	return feed
}

func getPostUrl(feeds []Feed) string {
	feedIndex, _ := feedsTable.GetSelection()
	postIndex, _ := postsTable.GetSelection()
	postUrl := feeds[feedIndex].items[postIndex].url

	return postUrl
}

package main

func getPostUrl(feeds []Feed) string {
	feedIndex, _ := feedsTable.GetSelection()
	postIndex, _ := postsTable.GetSelection()
	postUrl := feeds[feedIndex].items[postIndex].url

	return postUrl
}

func getFeedUrl() string {
	feedIndex, _ := feedsTable.GetSelection()
	cellRef := feedsTable.GetCell(feedIndex, 0).GetReference()
	feedUrl := cellRef.(Feed).url

	return feedUrl
}

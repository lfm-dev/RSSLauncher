package main

func getPostUrl(feeds []Feed) string {
	feedIndex, _ := feedsTable.GetSelection()
	postIndex, _ := postsTable.GetSelection()
	postUrl := feeds[feedIndex].items[postIndex].url

	return postUrl
}

func getFeedUrl(feeds []Feed) string {
	feedIndex, _ := feedsTable.GetSelection()
	postUrl := feeds[feedIndex].url

	return postUrl
}

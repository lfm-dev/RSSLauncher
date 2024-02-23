package main

func getPostUrl(feeds []Feed) string {
	feedIndex, _ := feedsTable.GetSelection()
	postIndex, _ := postsTable.GetSelection()
	postUrl := feeds[feedIndex].items[postIndex].url

	return postUrl
}

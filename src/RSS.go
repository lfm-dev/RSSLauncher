package main

//TODO run command with item url
//TODO open item on browser on enter

func main() {
	feedsUrls := getFeedsUrl()
	feeds := getFeeds(feedsUrls)
	view(feeds)
}

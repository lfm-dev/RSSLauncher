package main

import (
	"os"
)

//TODO run command with item url
//TODO open item on browser on enter

var (
	homePath, _   = os.UserHomeDir()
	feedsFilePath = homePath + "/.config/RSS/feeds.txt"
)

const (
	BROWSER = "firefox"
)

func main() {
	feedsUrls := getFeedsUrl()
	feeds := getFeeds(feedsUrls)
	view(feeds)
}

package main

import (
	"os"
)

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

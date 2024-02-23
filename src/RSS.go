package main

import (
	"os"
)

var (
	homePath, _    = os.UserHomeDir()
	feedsFilePath  = homePath + "/.config/RSS/feeds.txt"
	ignoreFilePath = homePath + "/.config/RSS/ignore.txt"
	wordsToIgnore  = getWordsToIgnore()
)

const (
	BROWSER = "firefox"
)

func main() {
	feeds := getFeeds()
	view(feeds)
}

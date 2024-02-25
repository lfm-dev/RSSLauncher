package main

import (
	"os"
)

var (
	homePath, _      = os.UserHomeDir()
	feedsFilePath    = homePath + "/.config/RSS/feeds.txt"
	ignoreFilePath   = homePath + "/.config/RSS/ignore.txt"
	commandsFilePath = homePath + "/.config/RSS/commands.csv"
	wordsToIgnore    = getWordsToIgnore()
	commands         = getCommands()
)

const (
	BROWSER = "firefox"
)

func main() {
	feeds := getFeeds()
	view(feeds)
}

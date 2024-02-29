package main

import (
	"os"
	"path"
)

//TODO cmd to add a new feed
//TODO cmd to clean cache

var (
	homePath, _      = os.UserHomeDir()
	configFolderPath = path.Join(homePath, "/.config/RSS")
	feedsFilePath    = path.Join(configFolderPath, "feeds.txt")
	ignoreFilePath   = path.Join(configFolderPath, "ignore.txt")
	commandsFilePath = path.Join(configFolderPath, "commands.csv")
	DBFilePath       = path.Join(configFolderPath, "oldItems.db")

	wordsToIgnore = getFileLines(ignoreFilePath)
	commands      = getCommands()
	DB            = openDB()
)

func main() {
	feeds := getFeeds()
	view(feeds)
}

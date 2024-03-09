package main

import (
	"os"
	"path"
)

//TODO cmd to add a new feed
//TODO cmd to clean cache
//TODO cmd to import opml
//TODO cmd to export settings, feeds, etc
//TODO copy url

var (
	homePath, _      = os.UserHomeDir()
	configFolderPath = path.Join(homePath, "/.config/RSSLauncher")
	feedsFilePath    = path.Join(configFolderPath, "feeds.txt")
	ignoreFilePath   = path.Join(configFolderPath, "ignore.txt")
	commandsFilePath = path.Join(configFolderPath, "commands.csv")
	DBFilePath       = path.Join(configFolderPath, "items.db")

	wordsToIgnore = getFileLines(ignoreFilePath)
	commands      = getCommands()
	DB            = openDB()
)

func main() {
	feeds := getFeeds()
	view(feeds)
}

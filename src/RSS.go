package main

import (
	"os"
	"path"
)

//TODO cmd to add a new feed (with category option)
//TODO cmd to import opml

var (
	homePath, _       = os.UserHomeDir()
	configFolderPath  = path.Join(homePath, "/.config/RSSLauncher")
	feedsFilePath     = path.Join(configFolderPath, "feeds.txt")
	blacklistFilePath = path.Join(configFolderPath, "blacklist.txt")
	commandsFilePath  = path.Join(configFolderPath, "commands.csv")
	DBFilePath        = path.Join(configFolderPath, "items.db")

	blacklistedWords = getFileLines(blacklistFilePath)
	commands         = getCommands()
	DB               = openDB()
)

func main() {
	runUsrInput()
	feeds := getFeeds()
	view(feeds)
}

package main

import (
	"os"
	"path"
)

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

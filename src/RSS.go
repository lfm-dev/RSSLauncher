package main

import (
	"os"
	"path"
)

var (
	homePath, _      = os.UserHomeDir()
	configFolderPath = path.Join(homePath, "/.config/RSS")
	feedsFilePath    = path.Join(configFolderPath, "feeds.txt")
	ignoreFilePath   = path.Join(configFolderPath, "ignore.txt")
	commandsFilePath = path.Join(configFolderPath, "commands.csv")
	DBFilePath       = path.Join(configFolderPath, "DB.csv")

	wordsToIgnore = getFileLines(ignoreFilePath)
	commands      = getCommands()
)

func main() {
	feeds := getFeeds()
	view(feeds)
}

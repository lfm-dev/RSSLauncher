package main

import (
	"errors"
	"os"
	"strings"
)

func getCommandList() map[string]string {
	commands := make(map[string]string)
	commandsFileContent, err := os.ReadFile(commandsFilePath)
	if err != nil {
		panic(errors.New("can't read commands.csv"))
	}
	commandsLines := strings.Split(strings.TrimSpace(string(commandsFileContent)), "\n")
	for _, line := range commandsLines {
		name, command := strings.Split(line, ",")[0], strings.Split(line, ",")[1]
		commands[name] = command
	}
	return commands
}

func getFeedsUrls() []string {
	feeds, err := os.ReadFile(feedsFilePath)
	if err != nil {
		panic(errors.New("can't read feeds.txt"))
	}
	feedsUrls := strings.Split(strings.TrimSpace(string(feeds)), "\n")
	return feedsUrls
}

func getWordsToIgnore() []string {
	ignoreFile, err := os.ReadFile(ignoreFilePath)
	if err != nil {
		return make([]string, 0)
	}
	wordsToIgnore := strings.Split(strings.TrimSpace(string(ignoreFile)), "\n")
	return wordsToIgnore
}

func itemHasIgnoredWord(title string) bool {
	for _, word := range wordsToIgnore {
		if strings.Contains(title, word) {
			return true
		}
	}
	return false
}

package main

import (
	"errors"
	"os"
	"strings"
)

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

func postHasIgnoredWord(title string) bool {
	for _, word := range wordsToIgnore {
		if strings.Contains(title, word) {
			return true
		}
	}
	return false
}

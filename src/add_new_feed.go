package main

import (
	"slices"
)

func addNewFeed(newFeedURL string) {
	feedFileLines := getFileLines(feedsFilePath)
	feedFileLines = slices.Insert(feedFileLines, 0, newFeedURL)
	writeLinesToFile(feedFileLines, feedsFilePath)
}

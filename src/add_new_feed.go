package main

import (
	"slices"
)

func addNewFeed(newFeedURL, newFeedCategory string) {
	feedFileLines := getFileLines(feedsFilePath)

	if newFeedCategory == "noCategory" {
		feedFileLines = slices.Insert(feedFileLines, 0, newFeedURL)

	} else if slices.Contains(feedFileLines, "#"+newFeedCategory) {
		i := slices.Index(feedFileLines, "#"+newFeedCategory)
		feedFileLines = slices.Insert(feedFileLines, i+1, newFeedURL)

	} else {
		feedFileLines = append(feedFileLines, "#"+newFeedCategory, newFeedURL)
	}

	writeLinesToFile(feedFileLines, feedsFilePath)
}

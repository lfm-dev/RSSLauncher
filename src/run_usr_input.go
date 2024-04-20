package main

import (
	"flag"
	"os"
)

func runUsrInput() {
	var clearCache = flag.Bool("cc", false, "clear cache")
	var exportSettings = flag.Bool("es", false, "export settings")
	var newFeedURL string
	flag.StringVar(&newFeedURL, "af", "", "add new feed")
	var newFeedCategory string
	flag.StringVar(&newFeedCategory, "c", "noCategory", "new feed category")
	var opmlFilePath string
	flag.StringVar(&opmlFilePath, "iopml", "", "import opml (opml path)")

	flag.Parse()

	if *clearCache {
		os.Remove(DBFilePath)
		os.Exit(0)
	}

	if *exportSettings {
		exportSettingsToZip()
		os.Exit(0)
	}

	if len(newFeedURL) > 0 {
		addNewFeed(newFeedURL, newFeedCategory)
		os.Exit(0)
	}

	if len(opmlFilePath) > 0 {
		importOpmlFeeds(opmlFilePath)
		os.Exit(0)
	}
}

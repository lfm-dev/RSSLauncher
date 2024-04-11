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
		addNewFeed(newFeedURL)
		os.Exit(0)
	}
}

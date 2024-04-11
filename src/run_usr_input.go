package main

import (
	"flag"
	"fmt"
	"os"
)

func runUsrInput() {
	var clearCache = flag.Bool("cc", false, "clear cache")
	var exportSettings = flag.Bool("es", false, "export settings")
	var newFeed string
	flag.StringVar(&newFeed, "af", "", "add new feed")

	flag.Parse()

	if *clearCache {
		os.Remove(DBFilePath)
		os.Exit(0)
	}

	if *exportSettings {
		exportSettingsToZip()
		os.Exit(0)
	}

	if len(newFeed) > 0 {
		fmt.Println(newFeed)
		os.Exit(0)
	}
}

package main

import (
	"flag"
	"os"
)

func runUsrInput() {
	var clearCache = flag.Bool("cc", false, "clear cache")
	var exportSettings = flag.Bool("es", false, "export settings")

	flag.Parse()

	if *clearCache {
		os.Remove(DBFilePath)
		os.Exit(0)
	}

	if *exportSettings {
		exportSettingsToZip()
		os.Exit(0)
	}
}

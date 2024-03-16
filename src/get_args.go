package main

import (
	"flag"
	"os"
)

func getArgs() {
	var clearCache = flag.Bool("cc", false, "clear cache")

	flag.Parse()

	if *clearCache {
		os.Remove(DBFilePath)
		os.Exit(0)
	}
}

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func addFilesToZip(filename string, zipw *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open %s: %s", filename, err)
	}
	defer file.Close()

	wr, err := zipw.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to add %s in zip: %s", filename, err)
	}

	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("failed to write %s to zip: %s", filename, err)
	}

	return nil
}

func exportSettingsToZip() {
	zipFile, err := os.OpenFile("RSSLauncher.zip", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic("Failed to write zip")
	}
	defer zipFile.Close()

	var files = []string{"commands.csv", "feeds.txt", "blacklist.txt", "items.db"}
	zipw := zip.NewWriter(zipFile)
	defer zipw.Close()

	os.Chdir(configFolderPath)
	for _, filename := range files {
		if err := addFilesToZip(filename, zipw); err != nil {
			fmt.Println(err)
		}
	}

}

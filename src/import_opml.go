package main

import (
	"encoding/xml"
	"os"
	"slices"
)

type opml struct {
	Body body
}

type body struct {
	XMLName  xml.Name  `xml:"body"`
	Outlines []outline `xml:"outline"`
}

type outline struct {
	XmlUrl string `xml:"xmlUrl,attr"`
}

func importOpmlFeeds(opmlFilePath string) {
	opml := opml{}
	xmlContent, _ := os.ReadFile(opmlFilePath)
	err := xml.Unmarshal(xmlContent, &opml)
	if err != nil {
		panic(err)
	}

	feedFileLines := getFileLines(feedsFilePath)
	for i, outline := range opml.Body.Outlines {
		feedFileLines = slices.Insert(feedFileLines, i, outline.XmlUrl)
	}
	writeLinesToFile(feedFileLines, feedsFilePath)
}

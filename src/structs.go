package main

import "time"

type Feed struct {
	url        string
	name       string
	lastUpdate time.Time
	items      []FeedItem
}

type FeedItem struct {
	url          string
	title        string
	date         time.Time
	dateFormated string
}

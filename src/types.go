package main

import (
	"time"
)

type Feed struct {
	feedUrl string
	url     string
	name    string
	items   []FeedItem
}

type FeedItem struct {
	url   string
	title string
	date  time.Time
	read  bool
}

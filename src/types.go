package main

type Feed struct {
	feedUrl string
	url     string
	name    string
	items   []FeedItem
}

type FeedItem struct {
	url          string
	title        string
	dateFormated string
	read         bool
}

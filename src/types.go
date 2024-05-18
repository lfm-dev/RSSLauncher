package main

import (
	"time"
)

type Feed struct {
	url      string
	name     string
	category string
	items    []FeedItem
}

type FeedItem struct {
	FeedUrl        string
	ItemPrimaryKey string `gorm:"primarykey"`
	ItemUrl        string
	Title          string
	Date           time.Time
	Read           bool
}

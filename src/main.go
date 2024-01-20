package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://calnewport.com/blog/feed")
	if err != nil {
		fmt.Println("oops")
	}
	fmt.Println(feed.Title)
}
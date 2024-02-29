package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func openDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DBFilePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&FeedItem{})
	return db
}

func addNewItemsToDB(items []FeedItem) {
	for _, newItem := range items {
		var items []FeedItem
		DB.Find(&items, "item_url = ?", newItem.ItemUrl)
		if len(items) > 0 { // that items already is in the DB
			break
		}
		DB.Create(&newItem)
	}
}

func getAllItemsFromDB(feedUrl string) []FeedItem {
	var items []FeedItem
	DB.Find(&items, "feed_url = ?", feedUrl)
	return items
}

func markAsReadInDB(itemUrl string) {
	var items []FeedItem
	DB.Model(&items).Where("item_url = ?", itemUrl).Update("read", true)
}

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
		DB.Find(&items, "item_primary_key = ?", newItem.ItemPrimaryKey)
		if len(items) == 0 { // item not found in DB
			DB.Create(&newItem)
		}
	}
}

func getAllItemsFromDB(feedUrl string) []FeedItem {
	var items []FeedItem
	DB.Order("date desc").Find(&items, "feed_url = ?", feedUrl)
	return items
}

func markAsReadInDB(itemUrl string) {
	var items []FeedItem
	DB.Model(&items).Where("item_url = ?", itemUrl).Update("read", true)
}

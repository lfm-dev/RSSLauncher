package main

import (
	"github.com/gdamore/tcell/v2"
)

func setupFeedsTable() {
	feedsTable.SetTitle("Feeds").SetBorder(true)

	feedsTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {

		case tcell.KeyRight, tcell.KeyTab:
			app.SetFocus(itemsTable)
			return nil

		case tcell.KeyEnter:
			feedUrl := getFeedData().url
			if len(feedUrl) > 0 { // only if feed has a web url
				go runCommand(feedUrl, commands["onEnter"])
			}
			return nil

		case tcell.KeyCtrlR:
			markAllFeedItemsAsRead()
			return nil

		case tcell.KeyCtrlX:
			copyToClipboard(getFeedData().url)
			return nil

		case tcell.KeyEscape:
			app.Stop()
			return nil

		default:
			return event
		}
	})

	feedsTable.SetSelectionChangedFunc(func(row int, _ int) {
		renderItemsTable(true)

		if row == 1 && getFeedData().category != "noCategory" {
			feedsTable.ScrollToBeginning() // so it shows the first category name
		}
	})
}

func setupItemsTable() {
	itemsTable.SetTitle("Items").SetBorder(true)

	itemsTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {

		case tcell.KeyLeft:
			app.SetFocus(feedsTable)
			return nil

		case tcell.KeyEnter:
			markItemAsRead()
			itemUrl := getItemData().ItemUrl
			go runCommand(itemUrl, commands["onEnter"])
			return nil

		case tcell.KeyCtrlR:
			markItemAsRead()
			return nil

		case tcell.KeyCtrlX:
			copyToClipboard(getItemData().ItemUrl)
			return nil

		case tcell.KeyTab:
			app.SetFocus(commandInput)
			return nil

		case tcell.KeyEscape:
			app.Stop()
			return nil

		default:
			return event
		}
	})
}

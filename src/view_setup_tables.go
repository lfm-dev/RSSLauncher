package main

import "github.com/gdamore/tcell/v2"

func setupFeedsTable() {
	feedsTable.SetTitle("Feeds").SetBorder(true)

	feedsTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		switch event.Key() {

		case tcell.KeyRight:
			app.SetFocus(itemsTable)
			return nil

		case tcell.KeyEnter:
			feedUrl := getFeedData().url
			if len(feedUrl) > 0 { // only if feed has a web url
				runCommand(feedUrl, commands["onEnter"])
			}
			return nil

		case tcell.KeyCtrlR:
			markAllItemsRead()
			return nil

		case tcell.KeyEscape:
			app.Stop()
			return nil

		default:
			return event
		}
	})

	feedsTable.SetSelectionChangedFunc(func(_ int, _ int) {
		renderItemsTable(true)
		itemsTable.ScrollToBeginning()
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
			runCommand(itemUrl, commands["onEnter"])
			return nil

		case tcell.KeyCtrlR:
			markItemAsRead()
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

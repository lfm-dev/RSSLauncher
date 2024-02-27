package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func setupFeedsTable() {
	feedsTable.SetTitle("Feeds")

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
	itemsTable.SetTitle("Items")

	itemsTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		switch event.Key() {

		case tcell.KeyLeft:
			app.SetFocus(feedsTable)
			return nil

		case tcell.KeyEnter:
			markItemAsRead()
			itemUrl := getItemData().url
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

func setupCommandInput() {
	commandInput.SetDoneFunc(func(key tcell.Key) {
		defer commandInput.SetText("")
		defer app.SetFocus(itemsTable)

		if key == tcell.KeyEnter && len(commandInput.GetText()) > 0 {
			markItemAsRead()
			itemUrl := getItemData().url
			if command, ok := commands[commandInput.GetText()]; ok {
				runCommand(itemUrl, command)
			} else {
				runCommand(itemUrl, commandInput.GetText()) // run custom command
			}
		}
	})
}

func setupUI(feeds []Feed) {
	tablesFlex.AddItem(feedsTable, 0, 1, false).AddItem(itemsTable, 0, 3, false)

	mainFlex.SetDirection(tview.FlexRow).AddItem(tablesFlex, 0, 1, false).AddItem(helpText, 1, 0, false).AddItem(commandList, 1, 0, false).AddItem(commandInput, 1, 0, false)

	feedsTable.SetBorder(true)
	itemsTable.SetBorder(true)

	renderFeedsTable(feeds)
	renderItemsTable(true)
}

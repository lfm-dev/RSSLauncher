package main

import (
	"github.com/rivo/tview"
)

var (
	app = tview.NewApplication()

	feedsTable = tview.NewTable().SetSelectable(true, false)
	itemsTable = tview.NewTable().SetSelectable(true, false)

	helpText = tview.NewTextView().SetText(
		"Open in browser: Enter / Run command: TAB / Mark as read: Ctrl+R / Quit: Esc",
	)
	commandList = tview.NewTextView().SetText(
		"Available commands: " + getCommandList(),
	)
	commandInput = tview.NewInputField().SetLabel("Command: ").SetFieldWidth(40).SetPlaceholder("use %url for custom commands")
)

func view(feeds []Feed) {
	mainFlex := tview.NewFlex()
	tablesFlex := tview.NewFlex()

	setupFeedsTable()
	setupItemsTable()
	setupCommandInput()

	tablesFlex.AddItem(feedsTable, 0, 1, false).AddItem(itemsTable, 0, 3, false)

	mainFlex.SetDirection(tview.FlexRow) // from top to bottom
	mainFlex.AddItem(tablesFlex, 0, 1, false).AddItem(helpText, 1, 0, false).AddItem(commandList, 1, 0, false).AddItem(commandInput, 1, 0, false)

	//render tables at startup
	renderFeedsTable(feeds)
	renderItemsTable(true)

	if err := app.SetRoot(mainFlex, true).SetFocus(feedsTable).Run(); err != nil {
		panic(err)
	}
}

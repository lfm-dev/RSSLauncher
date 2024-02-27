package main

import (
	"github.com/rivo/tview"
)

var (
	app = tview.NewApplication()

	feedsTable = tview.NewTable().SetSelectable(true, false)
	itemsTable = tview.NewTable().SetSelectable(true, false)

	commandInput = tview.NewInputField().SetLabel("Command: ").SetFieldWidth(40).SetPlaceholder("use %url for custom commands")
	helpText     = tview.NewTextView().SetText(
		"Open in browser: Enter / Run command: TAB / Mark as read: Ctrl+R / Quit: Esc",
	)
	commandList = tview.NewTextView().SetText(
		"Available commands: " + getCommandList(),
	)

	mainFlex   = tview.NewFlex()
	tablesFlex = tview.NewFlex()
)

func view(feeds []Feed) {
	setupFeedsTable()
	setupItemsTable()
	setupCommandInput()
	setupUI(feeds)

	if err := app.SetRoot(mainFlex, true).SetFocus(feedsTable).Run(); err != nil {
		panic(err)
	}

}

package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	unreadTextColor = tcell.ColorBlack //these colors suck
	unreadBkgColor  = tcell.ColorWhite
	readTextColor   = tcell.ColorGreen
	readBkgColor    = tcell.ColorYellow
)

var (
	app = tview.NewApplication()

	feedsTable = tview.NewTable().SetSelectable(true, false)
	itemsTable = tview.NewTable().SetSelectable(true, false)

	commandInput = tview.NewInputField().SetLabel("Command: ").SetFieldWidth(30)
	helpText     = tview.NewTextView().SetText(
		"Open in browser: Enter / Run command: TAB / Quit: Esc",
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

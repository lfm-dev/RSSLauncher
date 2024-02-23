package main

import (
	"github.com/rivo/tview"
)

var (
	app = tview.NewApplication()

	feedsTable = tview.NewTable().SetSelectable(true, false)
	postsTable = tview.NewTable().SetSelectable(true, false)

	commandInput = tview.NewInputField().SetLabel("Command: ").SetFieldWidth(30)
	helpText     = tview.NewTextView().SetText(
		"Open in browser: Enter / Run command: TAB / Quit: Esc",
	)

	mainFlex   = tview.NewFlex()
	tablesFlex = tview.NewFlex()
)

func view(feeds []Feed) {
	setupFeedsTable(feeds)
	setupPostsTable(feeds)
	setupCommandInput(feeds)
	setupUI(feeds)

	if err := app.SetRoot(mainFlex, true).SetFocus(feedsTable).Run(); err != nil {
		panic(err)
	}

}

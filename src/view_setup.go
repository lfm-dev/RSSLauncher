package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func setupFeedsTable(feeds []Feed) {
	feedsTable.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEscape:
			app.Stop()
		}
	})

	feedsTable.SetSelectedFunc(func(_ int, _ int) {
		app.SetFocus(postsTable)
	})

	feedsTable.SetSelectionChangedFunc(func(feedIndex int, _ int) {
		renderPostsTable(postsTable, feeds[feedIndex])
	})
}

func setupPostsTable(feeds []Feed) {
	postsTable.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyTab:
			app.SetFocus(commandInput)
		case tcell.KeyEscape:
			app.SetFocus(feedsTable)
		}
	})

	postsTable.SetSelectedFunc(func(_ int, _ int) {
		postUrl := getPostUrl(feeds)
		cmd := exec.Command(BROWSER, postUrl)
		cmd.Run()
		app.SetFocus(postsTable)
	})

	feedsTable.SetBorder(true)
	postsTable.SetBorder(true)
}

func setupCommandInput(feeds []Feed) {
	commandInput.SetDoneFunc(func(key tcell.Key) {
		defer commandInput.SetText("")
		defer app.SetFocus(postsTable)

		if key == tcell.KeyEnter && len(commandInput.GetText()) > 0 {
			postUrl := getPostUrl(feeds)
			command := strings.Split(
				strings.Replace(commandInput.GetText(), "%url", postUrl, 1),
				" ")

			cmd := exec.Command(command[0], command[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Run()
			app.Sync() // fix screen
		}
	})
}

func setupUI(feeds []Feed) {
	tablesFlex.AddItem(feedsTable, 0, 1, false).AddItem(postsTable, 0, 3, false)

	mainFlex.SetDirection(tview.FlexRow).AddItem(tablesFlex, 0, 1, false).AddItem(helpText, 1, 0, false).AddItem(commandInput, 1, 0, false)
	renderFeedsTable(feeds, feedsTable)
	renderPostsTable(postsTable, feeds[0]) // show first feed posts at start
}

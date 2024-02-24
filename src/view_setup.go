package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func setupFeedsTable() {
	feedsTable.SetTitle("Feeds")

	feedsTable.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEscape:
			app.Stop()
		}
	})

	feedsTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRight {
			app.SetFocus(postsTable)
			return nil
		}
		return event
	})

	feedsTable.SetSelectedFunc(func(_ int, _ int) {
		feedUrl := getFeedData().url

		if len(feedUrl) > 0 { // only if feed has a web url
			cmd := exec.Command(BROWSER, feedUrl)
			cmd.Run()
		}
	})

	feedsTable.SetSelectionChangedFunc(func(feedIndex int, _ int) {
		renderPostsTable()
		postsTable.ScrollToBeginning()
	})
}

func setupPostsTable() {
	postsTable.SetTitle("Posts")

	postsTable.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyTab:
			app.SetFocus(commandInput)
		case tcell.KeyEscape:
			app.Stop()
		}
	})

	postsTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyLeft {
			app.SetFocus(feedsTable)
			return nil
		}
		return event
	})

	postsTable.SetSelectedFunc(func(_ int, _ int) {
		postUrl := getPostData().url
		cmd := exec.Command(BROWSER, postUrl)
		cmd.Run()
	})

	feedsTable.SetBorder(true)
	postsTable.SetBorder(true)
}

func setupCommandInput() {
	commandInput.SetDoneFunc(func(key tcell.Key) {
		defer commandInput.SetText("")
		defer app.SetFocus(postsTable)

		if key == tcell.KeyEnter && len(commandInput.GetText()) > 0 {
			postUrl := getPostData().url
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

	renderFeedsTable(feeds)
	renderPostsTable()
}

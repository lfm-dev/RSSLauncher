package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	HELPTEXT = "Open in browser: Enter / Run command: TAB / Quit | Return: Esc"
)

var (
	app = tview.NewApplication()

	feedsTable = tview.NewTable().SetSelectable(true, false)
	postsTable = tview.NewTable().SetSelectable(true, false)

	commandInput = tview.NewInputField().SetLabel("Test: ").SetFieldWidth(30)
	helpText     = tview.NewTextView().SetText(HELPTEXT)

	mainFlex   = tview.NewFlex()
	tablesFlex = tview.NewFlex()
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

func getPostUrl(feeds []Feed) string {
	feedIndex, _ := feedsTable.GetSelection()
	postIndex, _ := postsTable.GetSelection()
	postUrl := feeds[feedIndex].items[postIndex].url

	return postUrl
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

			helpText.SetText("Running command...") //TODO fix

			cmd := exec.Command(command[0], command[1:]...)
			err := cmd.Run()

			if err != nil {
				helpText.SetText("ERROR")
			} else {
				helpText.SetText("Done!")
			}
		}
	})
}

func renderFeedsTable(feeds []Feed, feedsTable *tview.Table) {
	for i, feed := range feeds {
		feedsTable.SetCell(i, 0,
			tview.NewTableCell(feed.name).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
	}
}

func renderPostsTable(postsTable *tview.Table, feed Feed) {
	postsTable.Clear()
	for i, post := range feed.items {
		postLine := fmt.Sprintf("(%s) %s", post.dateFormated, post.title)
		postsTable.SetCell(i, 0,
			tview.NewTableCell(postLine))
	}
}

func view(feeds []Feed) {
	setupFeedsTable(feeds)
	setupPostsTable(feeds)
	setupCommandInput(feeds)

	tablesFlex.AddItem(feedsTable, 0, 1, false).AddItem(postsTable, 0, 3, false)
	mainFlex.SetDirection(tview.FlexRow)
	mainFlex.AddItem(tablesFlex, 0, 1, false).AddItem(helpText, 1, 0, false).AddItem(commandInput, 1, 0, false)

	renderFeedsTable(feeds, feedsTable)
	renderPostsTable(postsTable, feeds[0]) // show first feed posts at start

	if err := app.SetRoot(mainFlex, true).SetFocus(feedsTable).Run(); err != nil {
		panic(err)
	}

}

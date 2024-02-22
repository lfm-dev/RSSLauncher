package main

import (
	"fmt"
	"os/exec"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	mainFlex   = tview.NewFlex()
	tablesFlex = tview.NewFlex()
)

const (
	HELPTEXT = "Open in browser: Enter / Run command: TAB / Quit | Return: Esc"
)

func getTables(feeds []Feed, app *tview.Application, cmdInput *tview.InputField) (*tview.Table, *tview.Table) {
	feedsTable := tview.NewTable().SetSelectable(true, false)
	postsTable := tview.NewTable().SetSelectable(true, false)

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

	postsTable.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyTab:
			app.SetFocus(cmdInput)
		case tcell.KeyEscape:
			app.SetFocus(feedsTable)
		}
	})

	postsTable.SetSelectedFunc(func(itemIndex int, _ int) {
		feedIndex, _ := feedsTable.GetSelection()
		cmd := exec.Command(BROWSER, feeds[feedIndex].items[itemIndex].url)
		cmd.Run()
		app.SetFocus(postsTable)
	})

	feedsTable.SetBorder(true)
	postsTable.SetBorder(true)

	return feedsTable, postsTable
}

func getInputField(app *tview.Application) *tview.InputField {
	inputField := tview.NewInputField().SetLabel("Test: ").SetFieldWidth(30)

	inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter && len(inputField.GetText()) > 0 {
			cmd := exec.Command("firefox", inputField.GetText()) // TEST
			cmd.Run()
			inputField.SetText("")
			app.SetFocus(tablesFlex)
		}
		if key == tcell.KeyEscape {
			inputField.SetText("")
			app.SetFocus(tablesFlex)
		}
	})

	return inputField
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
	app := tview.NewApplication()
	cmdInput := getInputField(app)
	feedsTable, postsTable := getTables(feeds, app, cmdInput)
	helpText := tview.NewTextView().SetText(HELPTEXT)

	tablesFlex.AddItem(feedsTable, 0, 1, false).AddItem(postsTable, 0, 3, true) // postTable true so it is focused when press Esc in cmdInput
	mainFlex.SetDirection(tview.FlexRow)
	mainFlex.AddItem(tablesFlex, 0, 1, false).AddItem(helpText, 1, 0, false).AddItem(cmdInput, 1, 0, false)

	renderFeedsTable(feeds, feedsTable)
	renderPostsTable(postsTable, feeds[0]) // show first feed posts at start

	if err := app.SetRoot(mainFlex, true).SetFocus(feedsTable).Run(); err != nil {
		panic(err)
	}

}

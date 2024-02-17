package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func view(feeds []Feed) {

	app := tview.NewApplication()
	table := tview.NewTable().
		SetSelectable(true, false)

	var mainFlex = tview.NewFlex()
	var feedsFlex = tview.NewFlex()
	var postsFlex = tview.NewFlex()

	feedsFlex.AddItem(table, 0, 1, false)
	feedsFlex.SetBorder(true)
	mainFlex.AddItem(feedsFlex, 0, 1, false)

	rows := len(feeds)
	for row := 0; row < rows; row++ {
		table.SetCell(row, 0,
			tview.NewTableCell(feeds[row].name).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
	}

	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
	}).SetSelectedFunc(func(row int, column int) {
		// fmt.Println(table.GetCell(row, column).Text)
		postsFlex.AddItem(table, 0, 1, false)
		postsFlex.SetBorder(true)
		mainFlex.AddItem(postsFlex, 0, 2, false)
	})

	if err := app.SetRoot(mainFlex, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}

}

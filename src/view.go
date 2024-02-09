package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func view(feeds []Feed) {

	app := tview.NewApplication()
	table := tview.NewTable().
		SetSelectable(true, false)

	_, rows := 2, len(feeds)
	for row := 0; row < rows; row++ {
		table.SetCell(row, 0,
			tview.NewTableCell(feeds[row].name).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		// table.SetCell(row, 1,
		// 	tview.NewTableCell("").
		// 		SetTextColor(tcell.ColorWhite).
		// 		SetAlign(tview.AlignCenter))
	}

	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
	}).SetSelectedFunc(func(row int, column int) {
		fmt.Println(table.GetCell(row, column).Text)
	})

	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}

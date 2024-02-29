package main

import "github.com/gdamore/tcell/v2"

func setupCommandInput() {
	commandInput.SetDoneFunc(func(key tcell.Key) {
		defer commandInput.SetText("")
		defer app.SetFocus(itemsTable)

		if key == tcell.KeyEnter && len(commandInput.GetText()) > 0 {
			markItemAsRead()
			itemUrl := getItemData().ItemUrl
			if command, ok := commands[commandInput.GetText()]; ok {
				go runCommand(itemUrl, command)
			} else {
				go runCommand(itemUrl, commandInput.GetText()) // run custom command
			}
		}
	})
}

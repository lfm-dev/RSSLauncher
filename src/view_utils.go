package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/rivo/tview"
)

func getCommandList() string {
	var commandList []string
	for command := range commands {
		if command != "onEnter" {
			commandList = append(commandList, command)
		}
	}
	return strings.Join(commandList, " / ")
}

func getSelectedCell(table *tview.Table) *tview.TableCell {
	cellIndex, _ := table.GetSelection()

	return table.GetCell(cellIndex, 0)
}

func getFeedData() Feed {
	cellRef := getSelectedCell(feedsTable).GetReference()
	feed := cellRef.(Feed)

	return feed
}

func getItemData() FeedItem {
	cellRef := getSelectedCell(itemsTable).GetReference()
	item := cellRef.(FeedItem)

	return item
}

func runCommand(url string, command string) {
	cmd := strings.Split(
		strings.Replace(command, "%url", url, 1),
		" ")
	process := exec.Command(cmd[0], cmd[1:]...)
	process.Stderr = os.Stderr
	process.Stdout = os.Stdout
	process.Run()
}

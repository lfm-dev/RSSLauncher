package main

import (
	"os/exec"
	"strings"

	"github.com/rivo/tview"
	"golang.design/x/clipboard"
)

func getCommandList() string {
	var commandList []string
	for command := range commands {
		if command != "onEnter" {
			commandList = append(commandList, command)
		}
	}
	if len(commandList) == 0 {
		return "no available commands, configure them on commands.csv"
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

func feedHasUnreadItems(feed Feed) bool {
	for _, item := range feed.items {
		if !item.Read {
			return true
		}
	}
	return false
}

func runCommand(url string, command string) {
	cmd := strings.Split(
		strings.Replace(command, "%url", url, 1),
		" ")
	process := exec.Command(cmd[0], cmd[1:]...)
	process.Run()
	// app.Sync() // not longer really needed
}

func copyToClipboard(stringToCopy string) {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	clipboard.Write(clipboard.FmtText, []byte(stringToCopy))
}

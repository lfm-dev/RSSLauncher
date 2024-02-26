package main

import (
	"fmt"
	"os"
	"strings"
)

func getFileLines(filePath string) []string {
	isComment := func(s string) bool { return strings.HasPrefix(s, "//") }

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Errorf("can't read %s", filePath))
	}
	fileLines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")

	fileLinesNoComments := make([]string, 0)
	for _, line := range fileLines {
		if !isComment(line) {
			fileLinesNoComments = append(fileLinesNoComments, line)
		}
	}
	return fileLinesNoComments
}

func getCommands() map[string]string {
	commands := make(map[string]string)
	commandsLines := getFileLines(commandsFilePath)
	for _, line := range commandsLines {
		name, command := strings.Split(line, ",")[0], strings.Split(line, ",")[1]
		commands[name] = command
	}
	return commands
}

func itemHasIgnoredWord(title string) bool {
	for _, word := range wordsToIgnore {
		if strings.Contains(title, word) {
			return true
		}
	}
	return false
}

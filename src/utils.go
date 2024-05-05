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
		line = strings.TrimSpace(line)
		if !isComment(line) && len(line) > 0 {
			fileLinesNoComments = append(fileLinesNoComments, line)
		}
	}
	return fileLinesNoComments
}

func getCommands() map[string]string {
	commands := make(map[string]string)
	commandsLines := getFileLines(commandsFilePath)
	for _, line := range commandsLines {
		commandShortcut, command := strings.Split(line, ",")[0], strings.Split(line, ",")[1]
		commands[commandShortcut] = command
	}
	return commands
}

func itemHasBlacklistedWord(title string) bool {
	for _, blacklistedWord := range blacklistedWords {
		if strings.Contains(title, blacklistedWord) {
			return true
		}
	}
	return false
}

func writeLinesToFile(lines []string, filePath string) {
	if _, err := os.Stat(filePath); err == nil {
		os.Remove(filePath)
	}
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, line := range lines {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}
}

package utils

import (
	"os"
	"strings"
)

func ReadLines(filename string) []string {
	fileString := ReadText(filename)
	var listOfStrings []string
	for _, line := range strings.Split(strings.TrimSuffix(fileString, "\n"), "\n") {
		listOfStrings = append(listOfStrings, line)
	}

	return listOfStrings
}

func ReadText(filename string) string {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileString := string(fileData)

	return fileString
}

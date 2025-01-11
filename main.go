package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func readFileByLine(filePath string) []string {
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error: Enter a correct file path.")
		os.Exit(1)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}

func main() {
	// Check if file path is provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./dumb-grep <file_path> [search_pattern]")
		os.Exit(1)
	}

	filePath := os.Args[1]
	fileLines := readFileByLine(filePath)

	// Set default regex pattern to ".*" if the second argument is missing
	regexPattern := ".*"
	if len(os.Args) > 2 {
		regexPattern = os.Args[2]
	}

	for _, line := range fileLines {
		isPatternMatched, err := regexp.MatchString(regexPattern, line)
		if err != nil {
			fmt.Println("Error: Enter a valid regex.")
			os.Exit(1)
		}
		if isPatternMatched {
			fmt.Println(line)
		}
	}
}

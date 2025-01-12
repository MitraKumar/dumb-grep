package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/fatih/color"
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
	re, err := regexp.Compile(regexPattern)
	if err != nil {
		fmt.Println("Error: Enter a valid regex.")
		os.Exit(1)
	}

	for _, line := range fileLines {
		isPatternMatched := re.MatchString(line)
		if isPatternMatched {
			fmt.Println(re.ReplaceAllStringFunc(line, func(match string) string {
				return fmt.Sprintf("%s", color.GreenString(match))
			}))
		}
	}
}

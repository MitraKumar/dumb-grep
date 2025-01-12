package main

import (
	"fmt"
	"os"

	"github.com/MitraKumar/dumb-grep/lib"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./dumb-grep <file_path> [search_pattern]")
		os.Exit(1)
	}

	filePath := os.Args[1]
	regexPattern := ".*"
	if len(os.Args) > 2 {
		regexPattern = os.Args[2]
	}

	matchHighlighter, error := lib.NewMatchHighlighter(regexPattern)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	fileLines, error := matchHighlighter.HighlightFileLinesByPatter(filePath)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	for _, line := range fileLines {
		fmt.Println(line)
	}
}

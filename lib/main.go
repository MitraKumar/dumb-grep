package lib

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/fatih/color"
)

type MatchHighlighter struct {
	pattern *regexp.Regexp
}

func NewMatchHighlighter(pattern string) (*MatchHighlighter, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("error parsing regex: %w", err)
	}
	return &MatchHighlighter{pattern: re}, nil
}

func (m *MatchHighlighter) HighlightFileLinesByPatter(filePath string) ([]string, error) {
	var fileLines []string
	readFile, err := os.Open(filePath)
	if err != nil {
		return fileLines, fmt.Errorf("Something went wrong")
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if m.isPatternMatched(line) {
			fileLines = append(fileLines, m.highlightString(line))
		}
	}

	return fileLines, nil
}

func (m *MatchHighlighter) isPatternMatched(text string) bool {
	return m.pattern.MatchString(text)
}

func (m *MatchHighlighter) highlightString(text string) string {
	highlitedText := m.pattern.ReplaceAllStringFunc(text, func(match string) string {
		return fmt.Sprintf("%s", color.GreenString(match))
	})
	return highlitedText
}

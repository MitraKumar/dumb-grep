package lib

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/fatih/color"
)

type LineMatch struct {
	matchText  string
	lineNumber int
}

func newLineMatch(matchText string, lineNumber int) LineMatch {
	return LineMatch{
		matchText:  matchText,
		lineNumber: lineNumber,
	}
}

func (line *LineMatch) RenderLineMatch() {
	fmt.Printf("%d %s\n", line.lineNumber, line.matchText)
}

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

func (m *MatchHighlighter) HighlightFileLinesByPatter(filePath string) ([]LineMatch, error) {
	var fileLines []LineMatch
	readFile, err := os.Open(filePath)
	if err != nil {
		return fileLines, fmt.Errorf("something went wrong")
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	lineNumber := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lineNumber += 1
		if m.isPatternMatched(line) {
			fileLines = append(fileLines, newLineMatch(m.highlightString(line), lineNumber))
		}
	}

	return fileLines, nil
}

func (m *MatchHighlighter) isPatternMatched(text string) bool {
	return m.pattern.MatchString(text)
}

func (m *MatchHighlighter) highlightString(text string) string {
	highlitedText := m.pattern.ReplaceAllStringFunc(text, func(match string) string {
		return color.GreenString(match)
	})
	return highlitedText
}

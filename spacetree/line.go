package spacetree

import (
	"bufio"
	"strings"
)

// toLines process the scanner and generates lines
// with the value and the level on the tree
func toLines(scanner *bufio.Scanner, indentationSymbol string) []line {
	lines := make([]line, 0)

	for scanner.Scan() {
		text := scanner.Text()

		indentationLevel := calculateLevel(text, indentationSymbol)

		lines = append(lines, line{
			Value: strings.TrimLeft(text, indentationSymbol),
			Level: indentationLevel,
		})
	}

	return lines
}

// calculateLevel calculates the level of the line on the tree
// based on the indentation symbol occurences in the beggining
// of the line
func calculateLevel(text string, indentationSymbol string) int {
	level := 0
	idx := 0
	buf := text

	for strings.HasPrefix(buf, indentationSymbol) {
		level++
		idx += len(indentationSymbol)
		buf = text[idx:]
	}

	return level
}

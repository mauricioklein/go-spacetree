package spacetree

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toLines_OneLevelWithSpaces(t *testing.T) {
	scanner, closeFile := toScanner("one_level_with_spaces.txt")
	indentationSymbol := "  "

	defer closeFile()

	lines := toLines(scanner, indentationSymbol)

	assert.Equal(t, lines, []line{
		{Value: "1", Level: 0},
		{Value: "2", Level: 0},
		{Value: "3", Level: 0},
	})
}

func Test_toLines_TwoLevelsWithSpaces(t *testing.T) {
	scanner, closeFile := toScanner("two_levels_with_spaces.txt")
	indentationSymbol := "  "

	defer closeFile()

	lines := toLines(scanner, indentationSymbol)

	assert.Equal(t, lines, []line{
		{Value: "1", Level: 0},
		{Value: "1.1", Level: 1},
		{Value: "1.2", Level: 1},
		{Value: "2", Level: 0},
		{Value: "3", Level: 0},
		{Value: "3.1", Level: 1},
	})
}

func Test_toLines_MultipleLevelsWithSpaces(t *testing.T) {
	scanner, closeFile := toScanner("multiple_levels_with_spaces.txt")
	indentationSymbol := "  "

	defer closeFile()

	lines := toLines(scanner, indentationSymbol)

	assert.Equal(t, lines, []line{
		{Value: "1", Level: 0},
		{Value: "1.1", Level: 1},
		{Value: "1.1.1", Level: 2},
		{Value: "2", Level: 0},
		{Value: "2.1", Level: 1},
		{Value: "3", Level: 0},
		{Value: "3.1", Level: 1},
		{Value: "3.1.1", Level: 2},
	})
}

func Test_toLines_MultipleLevelsWithTabs(t *testing.T) {
	scanner, closeFile := toScanner("multiple_levels_with_tabs.txt")
	indentationSymbol := "\t"

	defer closeFile()

	lines := toLines(scanner, indentationSymbol)

	assert.Equal(t, lines, []line{
		{Value: "1", Level: 0},
		{Value: "1.1", Level: 1},
		{Value: "1.1.1", Level: 2},
		{Value: "2", Level: 0},
		{Value: "2.1", Level: 1},
		{Value: "3", Level: 0},
		{Value: "3.1", Level: 1},
		{Value: "3.1.1", Level: 2},
	})
}

func Test_toLines_MultipleLevelsWithDashes(t *testing.T) {
	scanner, closeFile := toScanner("multiple_levels_with_dashes.txt")
	indentationSymbol := "-"

	defer closeFile()

	lines := toLines(scanner, indentationSymbol)

	assert.Equal(t, lines, []line{
		{Value: "1", Level: 0},
		{Value: "1.1", Level: 1},
		{Value: "1.1.1", Level: 2},
		{Value: "2", Level: 0},
		{Value: "2.1", Level: 1},
		{Value: "3", Level: 0},
		{Value: "3.1", Level: 1},
		{Value: "3.1.1", Level: 2},
	})
}

func Test_toLines_IndentationBrokenWithSpaces(t *testing.T) {
	scanner, closeFile := toScanner("indentation_broken_with_spaces.txt")
	indentationSymbol := "  "

	defer closeFile()

	lines := toLines(scanner, indentationSymbol)

	assert.Equal(t, lines, []line{
		{Value: "1", Level: 0},
		{Value: "1.1", Level: 1},
		{Value: "1.2", Level: 1},
		{Value: "2", Level: 0},
		{Value: "3", Level: 0},
		{Value: "3.1", Level: 4}, // Indentation broken here!
	})
}

func toScanner(filename string) (*bufio.Scanner, func() error) {
	file, err := os.Open("../test_input/" + filename)

	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(file), file.Close
}

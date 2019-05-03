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

	assert.Equal(t, len(lines), 3)
	assert.Equal(t, lines[0], line{Value: "1", Level: 0})
	assert.Equal(t, lines[1], line{Value: "2", Level: 0})
	assert.Equal(t, lines[2], line{Value: "3", Level: 0})
}

func Test_toLines_TwoLevelsWithSpaces(t *testing.T) {
	scanner, closeFile := toScanner("two_levels_with_spaces.txt")
	indentationSymbol := "  "

	defer closeFile()

	lines := toLines(scanner, indentationSymbol)

	assert.Equal(t, len(lines), 6)
	assert.Equal(t, lines, []line{
		line{Value: "1", Level: 0},
		line{Value: "1.1", Level: 1},
		line{Value: "1.2", Level: 1},
		line{Value: "2", Level: 0},
		line{Value: "3", Level: 0},
		line{Value: "3.1", Level: 1},
	})
}

func Test_toLines_MultipleLevelsWithSpaces(t *testing.T) {
	scanner, closeFile := toScanner("multiple_levels_with_spaces.txt")
	indentationSymbol := "  "

	defer closeFile()

	lines := toLines(scanner, indentationSymbol)

	assert.Equal(t, len(lines), 8)
	assert.Equal(t, lines, []line{
		line{Value: "1", Level: 0},
		line{Value: "1.1", Level: 1},
		line{Value: "1.1.1", Level: 2},
		line{Value: "2", Level: 0},
		line{Value: "2.1", Level: 1},
		line{Value: "3", Level: 0},
		line{Value: "3.1", Level: 1},
		line{Value: "3.1.1", Level: 2},
	})
}

func Test_toLines_MultipleLevelsWithTabs(t *testing.T) {
	scanner, closeFile := toScanner("multiple_levels_with_tabs.txt")
	indentationSymbol := "\t"

	defer closeFile()

	lines := toLines(scanner, indentationSymbol)

	assert.Equal(t, len(lines), 8)
	assert.Equal(t, lines, []line{
		line{Value: "1", Level: 0},
		line{Value: "1.1", Level: 1},
		line{Value: "1.1.1", Level: 2},
		line{Value: "2", Level: 0},
		line{Value: "2.1", Level: 1},
		line{Value: "3", Level: 0},
		line{Value: "3.1", Level: 1},
		line{Value: "3.1.1", Level: 2},
	})
}

func Test_toLines_MultipleLevelsWithDashes(t *testing.T) {
	scanner, closeFile := toScanner("multiple_levels_with_dashes.txt")
	indentationSymbol := "-"

	defer closeFile()

	lines := toLines(scanner, indentationSymbol)

	assert.Equal(t, len(lines), 8)
	assert.Equal(t, lines, []line{
		line{Value: "1", Level: 0},
		line{Value: "1.1", Level: 1},
		line{Value: "1.1.1", Level: 2},
		line{Value: "2", Level: 0},
		line{Value: "2.1", Level: 1},
		line{Value: "3", Level: 0},
		line{Value: "3.1", Level: 1},
		line{Value: "3.1.1", Level: 2},
	})
}

func toScanner(filename string) (*bufio.Scanner, func() error) {
	file, err := os.Open("../test_input/" + filename)

	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(file), file.Close
}

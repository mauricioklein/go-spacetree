package spacetree

import (
	"fmt"
)

func newTree(lines []line) (*Node, error) {
	root := newNode("")

	levelNodeMapping := map[int]*Node{-1: root}
	level := -1

	for lineNumber, line := range lines {
		levelDiff := line.Level - level

		if levelDiff > 1 {
			return nil, newIndentationBrokenError(lineNumber)
		}

		parentNode := levelNodeMapping[line.Level-1]
		n := newNode(line.Value)

		parentNode.addChild(n)

		levelNodeMapping[line.Level] = n
		level = line.Level
	}

	return root, nil
}

func newIndentationBrokenError(lineNumber int) error {
	return fmt.Errorf("indentation broken at line %d", lineNumber)
}

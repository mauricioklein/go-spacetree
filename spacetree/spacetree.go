package spacetree

import "bufio"

// New returns the root node of the tree generated processing
// the provided scanner
func New(scanner *bufio.Scanner, indentationSymbol string) *Node {
	return newTree(
		toLines(scanner, indentationSymbol),
	)
}

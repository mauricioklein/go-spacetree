package spacetree

import (
	"github.com/golang-collections/collections/stack"
)

func newTree(lines []line) *Node {
	root := newNode("")
	lastNode := root

	stk := stack.New()
	stk.Push(root)

	currentLevel := -1

	for _, line := range lines {
		levelDiff := line.Level - currentLevel

		node := newNode(line.Value)

		if levelDiff > 1 {
			// Entry is a child of the last node
			stk.Push(lastNode)
			currentLevel++
		} else if levelDiff < 1 {
			// Entry is in a higher (or same level) than last node.
			// So, unstack by the level diff
			for i := 0; i < abs(levelDiff)+1; i++ {
				stk.Pop()
				currentLevel--
			}
		}

		// Add the new node to the tree
		stk.Peek().(*Node).AddChild(node)

		// Update the last node
		lastNode = node
	}

	return root
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

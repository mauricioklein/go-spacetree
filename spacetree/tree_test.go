package spacetree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newTree_OneLevelIndentation(t *testing.T) {
	lines := []line{
		{Value: "1", Level: 0},
		{Value: "2", Level: 0},
		{Value: "3", Level: 0},
	}

	root := newTree(lines)

	assert.Equal(t, root.Children[0], newNode("1"))
	assert.Equal(t, root.Children[1], newNode("2"))
	assert.Equal(t, root.Children[2], newNode("3"))
}

func Test_newTree_TwoLevelIndentation(t *testing.T) {
	lines := []line{
		{Value: "1", Level: 0},
		{Value: "1.1", Level: 1},
		{Value: "1.2", Level: 1},
		{Value: "2", Level: 0},
		{Value: "2.1", Level: 1},
	}

	root := newTree(lines)
	var n *Node

	// Node "1"
	n = root.Children[0]
	assert.Equal(t, n.Value, "1")
	assert.Equal(t, n.Children[0], newNode("1.1"))
	assert.Equal(t, n.Children[1], newNode("1.2"))

	// Node "2"
	n = root.Children[1]
	assert.Equal(t, n.Value, "2")
	assert.Equal(t, n.Children[0], newNode("2.1"))
}

func Test_newTree_MultiLevelIndentation(t *testing.T) {
	lines := []line{
		{Value: "1", Level: 0},
		{Value: "1.1", Level: 1},
		{Value: "1.1.1", Level: 2},
		{Value: "2", Level: 0},
		{Value: "3", Level: 0},
		{Value: "3.1", Level: 1},
	}

	root := newTree(lines)
	var n *Node

	// Node "1"
	n = root.Children[0]
	assert.Equal(t, n.Value, "1")
	assert.Equal(t, n.Children[0].Value, "1.1")

	// Node "1.1"
	n = root.Children[0].Children[0]
	assert.Equal(t, n.Value, "1.1")
	assert.Equal(t, n.Children[0], newNode("1.1.1"))

	// Node "2"
	n = root.Children[1]
	assert.Equal(t, n.Value, "2")

	// Node "3"
	n = root.Children[2]
	assert.Equal(t, n.Value, "3")
	assert.Equal(t, n.Children[0], newNode("3.1"))
}

func Test_newTree_Empty(t *testing.T) {
	lines := []line{}

	root := newTree(lines)

	assert.Equal(t, root, newNode(""))
}

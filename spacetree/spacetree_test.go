package spacetree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	scanner, closeFile := toScanner("one_level_with_spaces.txt")
	indentationSymbol := "  "

	defer closeFile()

	root, err := New(scanner, indentationSymbol)

	assert.NoError(t, err)

	assert.Equal(t, root, &Node{
		Value: "",
		Children: []*Node{
			newNode("1"),
			newNode("2"),
			newNode("3"),
		},
	})
}

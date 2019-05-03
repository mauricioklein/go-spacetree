package spacetree

// line represents a line with the respective
// value and and the level on the tree
type line struct {
	Value string
	Level int
}

// Node represents a node in the
// resulting tree
type Node struct {
	Value    string
	Children []*Node
}

// newNode creates a new node with the given
// value and no children
func newNode(value string) *Node {
	return &Node{
		Value:    value,
		Children: []*Node{},
	}
}

// AddChild adds a child to the node
func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

package binarytree

import "testing"

func TestListLeafs(t *testing.T) {
	node := Node{
		value: "y",
		right: &Node{value: "x"},
		left:  &Node{value: "x1", right: &Node{value: "x2"}},
	}

	if len(node.listLeafs()) != 2 {
		t.Errorf("Expected to found 2. Current: %d", node.countLeafs())
	}
}

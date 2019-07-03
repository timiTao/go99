package binarytree

import "testing"

func TestListLeafs(t *testing.T) {
	node := NodeInstance{
		value: "y",
		right: &NodeInstance{value: "x"},
		left:  &NodeInstance{value: "x1", right: &NodeInstance{value: "x2"}},
	}

	if len(node.listLeafs()) != 2 {
		t.Errorf("Expected to found 2. Current: %d", node.countLeafs())
	}
}

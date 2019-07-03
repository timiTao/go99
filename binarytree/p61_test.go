package binarytree

import "testing"

func TestCountLeafs(t *testing.T) {
	node := &NodeInstance{
		value: "y",
		right: &NodeInstance{value: "x"},
		left:  &NodeInstance{value: "x1", right: &NodeInstance{value: "x2"}},
	}

	if countLeafs(node) != 2 {
		t.Errorf("Expected to found 2. Current: %d", countLeafs(node))
	}
}

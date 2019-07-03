package binarytree

import "testing"

func TestCountOne(t *testing.T) {
	node := NodeInstance{value: "x"}

	if node.countNodes() != 1 {
		t.Errorf("Invalid count. Expected: 1. Current: %d", node.countNodes())
	}
}

func TestCountTree(t *testing.T) {
	node := NodeInstance{value: "x", right: &NodeInstance{value: "y"}, left: &NodeInstance{value: "z", left: &NodeInstance{value: "w"}}}

	if node.countNodes() != 4 {
		t.Errorf("Invalid count. Expected: 4. Current: %d", node.countNodes())
	}
}

package binarytree

import "testing"

func TestCountOne(t *testing.T) {
	node := Node{value: "x"}

	if node.countNodes() != 1 {
		t.Errorf("Invalid count. Expected: 1. Current: %d", node.countNodes())
	}
}

func TestCountTree(t *testing.T) {
	node := Node{value: "x", right: &Node{value: "y"}, left: &Node{value: "z", left: &Node{value: "w"}}}

	if node.countNodes() != 4 {
		t.Errorf("Invalid count. Expected: 4. Current: %d", node.countNodes())
	}
}

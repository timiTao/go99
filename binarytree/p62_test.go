package binarytree

import "testing"

func TestListInternalsNodes(t *testing.T) {
	node := Node{
		value: "y",
		right: &Node{value: "x"},
		left:  &Node{value: "x1", right: &Node{value: "x2", right: &Node{value: "x3"}}},
	}

	if len(node.listInternalsNodes()) != 3 {
		t.Errorf("Expected to found 3. Current: %d", len(node.listInternalsNodes()))
	}
}
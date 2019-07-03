package binarytree

import "testing"

func TestListInternalsNodes(t *testing.T) {
	node := NodeInstance{
		value: "y",
		right: &NodeInstance{value: "x"},
		left:  &NodeInstance{value: "x1", right: &NodeInstance{value: "x2", right: &NodeInstance{value: "x3"}}},
	}

	if len(node.listInternalsNodes()) != 3 {
		t.Errorf("Expected to found 3. Current: %d", len(node.listInternalsNodes()))
	}
}
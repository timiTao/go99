package binarytree

import "testing"

func TestRootListInternalsNodesAtLevel(t *testing.T) {
	node := &NodeInstance{
		value: "y",
		right: &NodeInstance{value: "x"},
		left:  &NodeInstance{value: "x1", right: &NodeInstance{value: "x2", right: &NodeInstance{value: "x3"}}},
	}

	if len(listInternalsNodesAtLevel(node, 1)) != 1 {
		t.Errorf("Expected to found 1. Current: %d", len(listInternalsNodesAtLevel(node, 1)))
	}
}

func TestOutOfIndexListInternalsNodesAtLevel(t *testing.T) {
	node := &NodeInstance{
		value: "y",
		right: &NodeInstance{value: "x"},
		left:  &NodeInstance{value: "x1", right: &NodeInstance{value: "x2", right: &NodeInstance{value: "x3"}}},
	}

	if len(listInternalsNodesAtLevel(node, 5)) != 0 {
		t.Errorf("Expected to found 0. Current: %d", len(listInternalsNodesAtLevel(node, 5)))
	}
}

func TestLevelListInternalsNodesAtLevel(t *testing.T) {
	node := &NodeInstance{
		value: "y",
		right: &NodeInstance{value: "x"},
		left:  &NodeInstance{value: "x1", right: &NodeInstance{value: "x2", right: &NodeInstance{value: "x3"}}},
	}

	if len(listInternalsNodesAtLevel(node, 2)) != 2 {
		t.Errorf("Expected to found 2. Current: %d", len(listInternalsNodesAtLevel(node, 2)))
	}
}

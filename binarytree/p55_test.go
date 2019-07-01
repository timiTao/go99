package binarytree

import (
	"strings"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	node := Node{
		value: "y",
		right: &Node{value: "x"},
		left:  &Node{value: "x1", right: &Node{value: "x2"}},
	}

	if node.isBalanced() != true {
		t.Error("This tree is should be balanced")
	}
}

func TestIsNotBalanced(t *testing.T) {
	node := Node{value: "y", left: &Node{value: "x1", right: &Node{value: "x2"}}}

	if node.isBalanced() != false {
		t.Error("This tree is should be not balanced")
	}
}

func TestEmptyBuildTreeVariations(t *testing.T) {
	list := buildTreeVariations(0)
	if len(list) != 0 {
		t.Errorf("Expected to find 0 solutions. Current: %d", len(list))
	}
}

func TestOneNodeBuildTreeVariations(t *testing.T) {
	list := buildTreeVariations(1)
	if len(list) != 1 {
		t.Errorf("Expected to find 1 solutions. Current: %d", len(list))
	}
}

func TestTwoBuildTreeVariations(t *testing.T) {
	list := buildTreeVariations(2)
	if len(list) != 2 {
		t.Errorf("Expected to find 2 solutions. Current: %d", len(list))
	}
}

func TestTreeBuildTreeVariations(t *testing.T) {
	list := buildTreeVariations(3)
	if len(list) != 5 {
		var list2 []string
		for _, x := range list {
			list2 = append(list2, x.print())
		}
		t.Errorf("Expected to find 5 solutions. Current: %d. Trees: %s", len(list), strings.Join(list2, "; "))
	}
}

func TestBuildBalancedTree(t *testing.T) {
	list := buildBalancedTree(5)
	if len(list) != 4 {
		var list2 []string
		for _, x := range list {
			list2 = append(list2, x.print())
		}
		t.Errorf("Expected to find 4 solutions. Current: %d. Trees: %s", len(list), strings.Join(list2, "; "))
	}
}

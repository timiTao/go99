package binarytree

import "fmt"

type Node interface {
	hasRight() bool
	getRight() Node
	hasLeft() bool
	getLeft() Node
	getValue() interface{}
}

type NodeInstance struct {
	value string
	right Node
	left  Node
}

func (n *NodeInstance) hasRight() bool {
	return n.right != nil
}
func (n *NodeInstance) getRight() Node {
	return n.right
}
func (n *NodeInstance) hasLeft() bool {
	return n.left != nil
}
func (n *NodeInstance) getLeft() Node {
	return n.left
}
func (n *NodeInstance) getValue() interface{} {
	return n.value
}

func countNodes(n Node) int {
	if n.hasLeft() && !n.hasRight() {
		return countNodes(n.getLeft()) + 1
	}

	if !n.hasLeft() && n.hasRight() {
		return countNodes(n.getRight()) + 1
	}

	if n.hasLeft() && n.hasRight() {
		return countNodes(n.getLeft()) + countNodes(n.getRight()) + 1
	}

	return 1
}

func printTree(n Node) string {
	if n.hasLeft() && !n.hasRight() {
		return fmt.Sprintf("%s[%s, nil]", n.getValue(), printTree(n.getLeft()))
	}

	if !n.hasLeft() && n.hasRight() {
		return fmt.Sprintf("%s[nil, %s]", n.getValue(), printTree(n.getRight()))
	}

	if n.hasLeft() && n.hasRight() {
		return fmt.Sprintf("%s[%s, %s]", n.getValue(), printTree(n.getLeft()), printTree(n.getRight()))
	}

	return n.getValue().(string)
}

func isBalanced(n Node) bool {
	if n.hasLeft() && !n.hasRight() {
		return countNodes(n.getLeft()) <= 1
	}

	if !n.hasLeft() && n.hasRight() {
		return countNodes(n.getRight()) <= 1
	}

	if n.hasLeft() && n.hasRight() {
		diff := countNodes(n.getLeft()) - countNodes(n.getRight())
		return diff >= -1 && diff <= 1
	}

	return true
}

func buildTreeVariations(howManyNodes int) []Node {
	var list []Node
	if howManyNodes <= 0 {
		return list
	}
	if howManyNodes == 1 {
		return append(list, &NodeInstance{value: "x"})
	}

	for i := 0; i < howManyNodes; i++ {
		leftOptions := buildTreeVariations(i)
		rightOptions := buildTreeVariations(howManyNodes - 1 - i)

		if len(leftOptions) != 0 {
			for x := 0; x < len(leftOptions); x++ {
				for y := 0; y < len(rightOptions); y++ {
					list = append(list, &NodeInstance{value: "x", left: leftOptions[x], right: rightOptions[y]})
					list = append(list, &NodeInstance{value: "x", right: rightOptions[y]})
				}

				list = append(list, &NodeInstance{value: "x", left: leftOptions[x]})
			}
		} else {
			for y := 0; y < len(rightOptions); y++ {
				list = append(list, &NodeInstance{value: "x", right: rightOptions[y]})
			}
		}
	}

	var list2 []Node
	for _, x := range list {
		if countNodes(x) == howManyNodes {
			list2 = append(list2, x)
		}
	}

	return list2
}

func buildBalancedTree(countNodes int) []Node {
	var list []Node
	for _, x := range buildTreeVariations(countNodes) {
		if isBalanced(x) {
			list = append(list, x)
		}
	}

	return list
}

/* P61. Count the leaves of a binary tree */
func countLeafs(n Node) int {
	if n.hasLeft() && !n.hasRight() {
		return countLeafs(n.getLeft())
	}

	if !n.hasLeft() && n.hasRight() {
		return countLeafs(n.getRight())
	}

	if n.hasLeft() && n.hasRight() {
		return countLeafs(n.getLeft()) + countLeafs(n.getRight())
	}

	return 1
}

/* P61A. Collect the leaves of a binary tree in a list */
func listLeafs(n Node) []Node {
	if n.hasLeft() && !n.hasRight() {
		return listLeafs(n.getLeft())
	}

	if !n.hasLeft() && n.hasRight() {
		return listLeafs(n.getRight())
	}

	if n.hasLeft() && n.hasRight() {
		return append(listLeafs(n.getLeft()), listLeafs(n.getRight())...)
	}

	return []Node{n}
}

/* P62. Collect the internal nodes of a binary tree in a list */
func listInternalsNodes(n Node) []Node {
	if n.hasLeft() && !n.hasRight() {
		return append([]Node{n}, listInternalsNodes(n.getLeft())...)
	}

	if !n.hasLeft() && n.hasRight() {
		return append([]Node{n}, listInternalsNodes(n.getRight())...)
	}

	if n.hasLeft() && n.hasRight() {
		return append([]Node{n}, append(listInternalsNodes(n.getLeft()), listInternalsNodes(n.getRight())...)...)
	}

	return []Node{}
}

/* P62B. Collect the nodes at a given level in a list */
func listInternalsNodesAtLevel(n Node, level int) []Node {
	if level == 1 {
		return []Node{n}
	}
	if n.hasLeft() && !n.hasRight() {
		return listInternalsNodesAtLevel(n.getLeft(), level-1)
	}

	if !n.hasLeft() && n.hasRight() {
		return listInternalsNodesAtLevel(n.getRight(), level-1)
	}

	if n.hasLeft() && n.hasRight() {
		return append(listInternalsNodesAtLevel(n.getLeft(), level-1), listInternalsNodesAtLevel(n.getRight(), level-1)...)
	}

	return []Node{}
}

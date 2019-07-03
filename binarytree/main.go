package binarytree

import "fmt"

type Node interface {
	hasRight() bool
	getRight() *Node
	hasLeft() bool
	getLeft() *Node
	getValue() interface{}
}

type NodeInstance struct {
	value string
	right *NodeInstance
	left  *NodeInstance
}

func (n *NodeInstance) hasRight() bool {
	return n.right != nil
}
func (n *NodeInstance) getRight() *NodeInstance {
	return n.right
}
func (n *NodeInstance) hasLeft() bool {
	return n.left != nil
}
func (n *NodeInstance) getLeft() *NodeInstance {
	return n.left
}

func (n NodeInstance) countNodes() int {
	if n.hasLeft() && !n.hasRight() {
		return n.left.countNodes() + 1
	}

	if !n.hasLeft() && n.hasRight() {
		return n.right.countNodes() + 1
	}

	if n.hasLeft() && n.hasRight() {
		return n.left.countNodes() + n.right.countNodes() + 1
	}

	return 1
}

func (n NodeInstance) print() string {
	if n.hasLeft() && !n.hasRight() {
		return fmt.Sprintf("%s[%s, nil]", n.value, n.left.print())
	}

	if !n.hasLeft() && n.hasRight() {
		return fmt.Sprintf("%s[nil, %s]", n.value, n.right.print())
	}

	if n.hasLeft() && n.hasRight() {
		return fmt.Sprintf("%s[%s, %s]", n.value, n.left.print(), n.right.print())
	}

	return n.value
}

func (n *NodeInstance) isBalanced() bool {
	if n.hasLeft() && !n.hasRight() {
		return n.left.countNodes() <= 1
	}

	if !n.hasLeft() && n.hasRight() {
		return n.right.countNodes() <= 1
	}

	if n.hasLeft() && n.hasRight() {
		diff := n.left.countNodes() - n.right.countNodes()
		return diff >= -1 && diff <= 1
	}

	return true
}

func buildTreeVariations(countNodes int) []*NodeInstance {
	var list []*NodeInstance
	if countNodes <= 0 {
		return list
	}
	if countNodes == 1 {
		return append(list, &NodeInstance{value: "x"})
	}

	for i := 0; i < countNodes; i++ {
		leftOptions := buildTreeVariations(i)
		rightOptions := buildTreeVariations(countNodes - 1 - i)

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

	list2 := []*NodeInstance{}
	for _, x := range list {
		if x.countNodes() == countNodes {
			list2 = append(list2, x)
		}
	}

	return list2
}

func buildBalancedTree(countNodes int) []*NodeInstance {
	list := []*NodeInstance{}
	for _, x := range buildTreeVariations(countNodes) {
		if x.isBalanced() {
			list = append(list, x)
		}
	}

	return list
}

/* P61. Count the leaves of a binary tree */

func (n *NodeInstance) countLeafs() int {
	if n.hasLeft() && !n.hasRight() {
		return n.left.countLeafs()
	}

	if !n.hasLeft() && n.hasRight() {
		return n.right.countLeafs()
	}

	if n.hasLeft() && n.hasRight() {
		return n.left.countLeafs() + n.right.countLeafs()
	}

	return 1
}

/* P61A. Collect the leaves of a binary tree in a list */
func (n *NodeInstance) listLeafs() []*NodeInstance {
	if n.hasLeft() && !n.hasRight() {
		return n.left.listLeafs()
	}

	if !n.hasLeft() && n.hasRight() {
		return n.right.listLeafs()
	}

	if n.hasLeft() && n.hasRight() {
		return append(n.left.listLeafs(), n.right.listLeafs()...)
	}

	return []*NodeInstance{n}
}

/* P62. Collect the internal nodes of a binary tree in a list */
func (n *NodeInstance) listInternalsNodes() []*NodeInstance {
	if n.hasLeft() && !n.hasRight() {
		return append([]*NodeInstance{n}, n.left.listInternalsNodes()...)
	}

	if !n.hasLeft() && n.hasRight() {
		return append([]*NodeInstance{n}, n.right.listInternalsNodes()...)
	}

	if n.hasLeft() && n.hasRight() {
		return append([]*NodeInstance{n}, append(n.left.listInternalsNodes(), n.right.listInternalsNodes()...)...)
	}

	return []*NodeInstance{}
}

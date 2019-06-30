package binarytree

import "fmt"

type Node struct {
	value string
	right *Node
	left  *Node
}

func (n *Node) hasRight() bool {
	return n.right != nil
}
func (n *Node) hasLeft() bool {
	return n.left != nil
}
func (n *Node) clone() *Node {
	return &Node{n.value, n.right.clone(), n.left.clone()}
}

func (n *Node) countNodes() int {
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

func (n *Node) print() string {
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

func (n *Node) isBalanced() bool {
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

func buildTreeVariations(countNodes int) []*Node {
	list := []*Node{}
	if countNodes <= 0 {
		return list
	}
	if countNodes == 1 {
		return append(list, &Node{value: "x"})
	}

	for i := 0; i < countNodes; i++ {
		leftOptions := buildTreeVariations(i)
		rightOptions := buildTreeVariations(countNodes - 1 - i)

		if len(leftOptions) != 0 {
			for x := 0; x < len(leftOptions); x++ {
				for y := 0; y < len(rightOptions); y++ {
					list = append(list, &Node{value: "x", left: leftOptions[x], right: rightOptions[y]})
					list = append(list, &Node{value: "x", right: rightOptions[y]})
				}

				list = append(list, &Node{value: "x", left: leftOptions[x]})
			}
		} else {
			for y := 0; y < len(rightOptions); y++ {
				list = append(list, &Node{value: "x", right: rightOptions[y]})
			}
		}
	}

	list2 := []*Node{}
	for _, x := range list {
		if x.countNodes() == countNodes {
			list2 = append(list2, x)
		}
	}

	return list2
}

func buildBalancedTree(countNodes int) []*Node {
	list := []*Node{}
	for _, x := range buildTreeVariations(countNodes) {
		if x.isBalanced() {
			list = append(list, x)
		}
	}

	return list
}

package binarytree

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

func (n *Node) countNodes() int {
	if (n.hasLeft() && !n.hasRight()) {
		return n.left.countNodes() + 1
	}

	if (!n.hasLeft() && n.hasRight()) {
		return n.right.countNodes() + 1
	}

	if (n.hasLeft() && n.hasRight()) {
		return n.left.countNodes() + n.right.countNodes() + 1
	}

	return 1
}

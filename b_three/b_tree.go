package b_tree

import (
	"fmt"
	"io"
)

type binaryNode struct {
	left  *binaryNode
	right *binaryNode
	data  rune
}

type BinaryTree struct {
	root *binaryNode
}

func (t *BinaryTree) Insert(data rune) *BinaryTree {
	if t.root == nil {
		t.root = &binaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *binaryNode) insert(data rune) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &binaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &binaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func print(w io.Writer, node *binaryNode, ns int, ch rune) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func Print(w io.Writer, t *BinaryTree) {
	if t == nil {
		return
	}
	if t.root == nil {
		return
	}

	fmt.Fprintf(w, "ROOT:%v\n", t.root.data)
	print(w, t.root.left, 6, 'L')
	print(w, t.root.right, 6, 'R')
}

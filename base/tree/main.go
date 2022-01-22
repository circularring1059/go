package main

import (
	"fmt"
)

type node struct {
	elem  interface{}
	left  *node
	right *node
}
type tree struct {
	root *node
}

func newTree() *tree {
	return &tree{nil}
}

func (t *tree) add(elem interface{}) {
	child := &node{elem, nil, nil}
	if t.root == nil {
		t.root = child
		return
	}
	queue := []*node{t.root}

	for {
		if len(queue) <= 0 {
			break
		}
		cur := queue[0]
		queue = append(queue[:0], queue[1:]...)
		if cur.left == nil {

			cur.left = child
			break
		} else {
			queue = append(queue, cur.left)
		}
		if cur.right == nil {
			cur.right = child
			break
		} else {
			queue = append(queue, cur.right)
		}

	}
}

func (t *tree) breadth_travel() {
	if t.root == nil {
		return
	}
	queue := []*node{t.root}
	for {
		if len(queue) <= 0 {
			break
		}
		cur := queue[0]
		queue = append(queue[:0], queue[1:]...)
		fmt.Println(cur.elem)
		if cur.left != nil {
			queue = append(queue, cur.left)
		}
		if cur.right != nil {
			queue = append(queue, cur.right)

		}
	}

}
func beforeInner(n *node) {
	if n == nil {
		return
	}
	fmt.Println(n.elem)
	beforeInner(n.left)
	beforeInner(n.right)
}

func (t *tree) beforeOrder() {
	if t.root == nil {
		return
	}
	cur := t.root
	beforeInner(cur)

}

func InInner(n *node) {
	if n == nil {
		return
	}
	beforeInner(n.left)
	fmt.Println(n.elem)
	beforeInner(n.right)
}

func (t *tree) InOrder() {
	if t.root == nil {
		return
	}
	cur := t.root
	InInner(cur)
}

func afterInner(n *node) {
	if n == nil {
		return
	}
	afterInner(n.left)
	afterInner(n.right)
	fmt.Println(n.elem)
}

func (t *tree) afterOrder() {
	if t.root == nil {
		return
	}
	cur := t.root
	afterInner(cur)
}
func main() {
	tree := newTree()
	tree.add(1)
	tree.add(2)
	tree.add(3)
	tree.add(4)
	// tree.breadth_travel()
	// tree.beforeOrder()
	// tree.InOrder()
	tree.afterOrder()

}

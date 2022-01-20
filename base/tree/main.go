package main

import "fmt"

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
	for _, cur := range queue {
		// cur := queue[0]
		fmt.Println(cur)
		queue = append(queue[:0], queue[1:]...)
		if cur.left == nil {
			cur.left = child
			return
		} else {
			queue = append(queue, cur.left)
		}
		if cur.right == nil {
			cur.right = child
			return
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
	for _, v := range queue {
		queue = append(queue[:0], queue[1:]...)
		fmt.Println(v.elem, "8")
		if v.left != nil {
			queue = append(queue, v.left)
		}
		if v.right != nil {
			queue = append(queue, v.right)

		}
	}

}

func main() {
	fmt.Println("hello")
	//newTree
	tree := newTree()
	// fmt.Println(tree)
	tree.add(1)
	tree.add(2)
	tree.add(3)
	tree.add(4)
	tree.breadth_travel()

}

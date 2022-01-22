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

func main() {
	tree := newTree()
	// fmt.Println(tree)
	tree.add(1)
	tree.add(2)
	tree.add(3)
	tree.add(4)
	tree.breadth_travel()

	// sl := []int{1, 2, 3, 4}
	// fmt.Println(sl)
	// for {
	// 	if len(sl) > 0 {
	// 		fmt.Println(sl[0])
	// 		sl = append(sl[:0], sl[1:]...)
	// 	} else {
	// 		break
	// 	}
	// }

}

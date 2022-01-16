package main

import (
	"fmt"
)

type node struct {
	data interface{}
	node *node
}

type link struct {
	node *node
}

func newLink() *link {
	return &link{nil}
}

// func (node *node) add(data interface{}) {
// 	new_node := node{12, nil}
// 	node.node = &new_node

// }

// func (node *node) is_empty() bool {
// 	if node.node == nil {
// 		return true
// 	}
// 	return false
// }

func main() {

	link := newLink()
	fmt.Println(link)

}

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

func (link *link) is_empty() bool {
	if link.node != nil {
		return false
	}
	return true
}

func (link *link) addNode(data interface{}) {
	newNode := &node{data, nil}
	if !link.is_empty() {
		cur := link.node
		newNode.node = cur
		link.node = newNode
	}
	link.node = newNode
}

func (link *link) append(data interface{}) {
	newNode := &node{data, nil}
	if !link.is_empty() {
		cur := link.node
		for cur.node != nil {
			cur = cur.node
		}
		cur.node = newNode
	} else {
		link.node = newNode
	}
}

func (link *link) length() int {
	if link.is_empty() {
		return 0
	}
	count := 1
	cur := link.node
	for cur.node != nil {
		cur = cur.node
		count++
	}
	return count

}

// func (link *link) insert(index int, data interface{}) {
// 	newNode := &node{data, nil}
// 	if !link.is_empty() {
// 		cur := link.node
// 		for cur.node != nil {
// 			cur = cur.node
// 		}
// 		cur.node = newNode
// 	} else {
// 		link.node = newNode
// 	}
// }

func (link *link) travel() {
	cur := link.node
	if cur != nil {
		for cur.node != nil {
			fmt.Println(cur.data)
			cur = cur.node
		}
		fmt.Println(cur.data)
	}
}

func main() {

	link := newLink()
	link.addNode("ring")
	link.append("file")
	link.append("dir")
	link.addNode("yuan")
	link.addNode("yuan")
	link.travel()
	fmt.Println(link.length())
}

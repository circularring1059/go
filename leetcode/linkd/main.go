package main

import (
	"fmt"
)

type Node struct {
	val  interface{}
	node *Node
}
type Link struct {
	head *Node
}

func main() {
	fmt.Println("start job")
	var link Link
	// fmt.Println(link) {<nil>}
	link.append(1)
	link.append(2)
	link.append(3)
	link.append(4)
	link.travel()
	fmt.Println("link lenght :", link.length())
}

func newNode(val interface{}) *Node {
	return &Node{
		val:  val,
		node: nil,
	}
}

func (L *Link) append(val interface{}) {
	new_node := newNode(val)
	if L.head == nil {
		L.head = new_node
	} else {
		cur := L.head
		for cur.node != nil {
			cur = cur.node
		}
		cur.node = new_node
	}
}

func (L *Link) isEmpty() bool {
	if L.head == nil {
		return true
	} else {
		return false
	}
}

func (L *Link) length() int {
	if L.isEmpty() {
		return 0
	} else {
		cur := L.head
		count := 0
		for cur != nil {
			count++
			cur = cur.node
		}
		return count
	}
}

func (L *Link) travel() {
	cur := L.head
	for cur != nil {
		fmt.Println(cur.val)
		cur = cur.node
	}
}


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
	link.insert(6, 9)
	link.travel()
	link.reverse1()
	link.travel()
	// fmt.Println("link lenght :", link.length())
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

func (L *Link) insert(val interface{}, index int) {
	if index < 0 {
		fmt.Println("worring: index is lt 0 setdefault 0")
		index = 0
	}
	if index > L.length()-1 {
		fmt.Println("worring:", val, "out of index setdefautl max_index:", L.length())
		index = L.length() - 1
	}
	new_node := newNode(val)
	if L.head == nil {
		L.head = new_node
	} else if index == L.length()-1 {
		L.append(val)
	} else if index == 0 {
		new_node.node = L.head
		L.head = new_node
	} else {
		cur := L.head
		for count := 1; count < index; count++ {
			cur = cur.node
		}
		new_node.node = cur.node
		cur.node = new_node
	}
}

func (L *Link) reverse() {
	var pre *Node = nil
	cur := L.head
	for cur != nil {
		next_node := cur.node
		cur.node = pre
		pre = cur
		cur = next_node
	}
	L.head = pre
}

func (L *Link) reverse1() {
	if L.head == nil || L.head.node == nil {
		return
	}
	cur := L.head
	cur = backtack(cur)
	L.head = cur
}

func backtack(node *Node) *Node {
	// node := N
	if node.node == nil {
		return node
	}
	newNode := backtack(node.node)
	node.node.node = node
	node.node = nil
	return newNode
}

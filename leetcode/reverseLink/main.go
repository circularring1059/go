package main

import (
	"fmt"

)

func main() {
	fmt.Println("start job")
	node1 := newNode(1)
	node2 := newNode(2)
	node3 := newNode(3)
	node4 := newNode(4)
	node1.node = node2
	node2.node = node3
	node3.node = node4
	node1.travel()
	new_node := reverseLink1(node1)
	node4.travel()
	new_node.travel()
	reverseLink(&new_node)
	node1.travel()

}

type Node struct {
	val  interface{}
	node *Node
}

func newNode(i interface{}) *Node {
	return &Node{
		val:  i,
		node: nil,
	}
}

func (n *Node) travel() {
	node := n
	for node != nil {

		if node.node == nil {
			fmt.Println(node.val, "-> nil")
			return
		}
		fmt.Print(node.val, "-> ")
		node = node.node
	}
}

func reverseLink(n *Node) Node {
	if n.node == nil {
		return *n
	} else {
		//进行reverse 后进行遍历会多出这个最后面node的 0 值
		// pre := &Node{
		// 	val: 0,
		//     node: nil,
		// }
		var pre *Node = nil
		for n != nil {
			next_node := n.node
			n.node = pre
			pre = n
			n = next_node
		}
		return *pre
	}
}

func reverseLink1(n *Node) Node {
	node := n
	if node.node == nil {
		return *node
	}
	node = backtack(node)
	return *node
}

func backtack(n *Node) *Node {
	//到最后一个node
	node := n
	if node.node == nil {
		return node
	} else {
		newNode := backtack(node.node)
		node.node.node = node
		node.node = nil
		return newNode
	}
}

package main

import "fmt"

type Node struct {
	value int
	node  *Node
}

func main() {
	fmt.Println("start job")
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)
	// node5 := NewNode(5)
	node1.node = node2
	node2.node = node3
	node3.node = node4
	// node4.node = node5
	// node3.node = node2
	fmt.Println(isRing(node1))
}

func NewNode(val int) *Node {
	return &Node{
		val,
		nil,
	}
}

func isRing(n *Node) bool {
	if n.node == nil || n.node.node == nil {
		return false
	}
	pre_point := n
	aft_point := n
	for aft_point != nil && aft_point.node != nil {
		pre_point = pre_point.node
		aft_point = aft_point.node.node
		if aft_point == nil {
			return false
		}
		if pre_point == aft_point || pre_point == aft_point.node {
			return true
		}
	}
	return false
}

package main 

import "fmt"

func main () {
	fmt.Println("start job")
	node1 := newNode(8)
	fmt.Println(node1)
}

type Node struct {
	val interface{}
	node *Node
}


func newNode(i interface{}) *Node{
	return &Node{
		val i,
		node nil,
	}
}
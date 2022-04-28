package main 

import "fmt"

func main () {
	fmt.Println("start job")
	node1 := newNode(1)
	node2 := newNode(2)
	node3 := newNode(3)
	node1.node = node2
	node2.node = node3
	node1.travel()
	reverLink(node1)
	node3.travel()
}

type Node struct {
	val interface{}
	node *Node
}


func newNode(i interface{}) *Node{
	return &Node{
		val: i,
		node: nil,
	}
}

func (n *Node) travel(){
	node := n
	for node != nil{
		fmt.Println(node.val)
		node = node.node
	}
}

func reverLink(n *Node) Node{
	if n.node == nil{
		return *n
	}else{
		//进行reverse 后进行遍历会多出这个最后面node的 0 值
		pre := &Node{
			val: 0,
		    node: nil,
		}
		for n != nil{
			next_node := n.node
			n.node = pre
			pre = n
			n = next_node
		}
		return *pre
	}
}
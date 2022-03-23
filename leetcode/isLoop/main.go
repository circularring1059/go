package ring

import "fmt"

type node struct {
	elem interface{}
	node *node
}

func isLoop(node *node) bool {
	if node == nil || node.node == nil {
		return false
	}
	fast := node
	slow := node
	for {
		if fast == nil || fast.node == nil {
			return false
		} else {
			fast = fast.node.node
			slow = slow.node
			if fast == nil {
				return false
			}
			if fast.node == slow || slow == fast {
				return true
			}
		}
	}
}

func main() {
	node1 := &node{1, nil}
	node2 := &node{2, node1}
	node3 := &node{3, node2}
	node4 := &node{4, node3}
	node1.node = node2

	fmt.Println(isLoop(node4))
}

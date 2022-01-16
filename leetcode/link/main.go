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

func (link *link) insert(index int, data interface{}) {
	if index < 0 || index >= link.length() {
		panic("out of index")
	}
	newNode := &node{data, nil}
	if link.is_empty() {
		link.node = newNode
	} else {
		if index == 0 {
			link.addNode(data)
		} else {
			cur := link.node
			for i := 1; i < index; i++ {
				cur = cur.node
			}
			newNode.node = cur.node
			cur.node = newNode
		}
	}
}

func (link *link) remove(index int) {
	if index < 0 || index >= link.length() {
		panic("out of index")
	}
	if link.is_empty() {
		panic("empty unsupport remove")
	}
	if index == 0 {
		if link.length() == 1 {
			link.node = nil
		} else {
			cur := link.node
			link.node = cur.node
		}
	} else {
		cur := link.node
		for i := 1; i < index; i++ {
			cur = cur.node
		}
		// cur_next := cur.node
		// cur.node = cur_next.node
		cur.node = cur.node.node
	}

}
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
	link.addNode("yuan")
	link.addNode("ring")
	link.append("file")
	link.append("dir")
	link.remove(3)
	// link.insert(1, 2)
	link.travel()
	fmt.Println(link.length())
}

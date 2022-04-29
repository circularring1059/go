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
	fmt.Println("length:", link.length())
	link.insert(6, 9)
	link.travel()
	link.reverse1()
	link.remove(4)
	link.remove(2)
	fmt.Println(link.search(1))
	link.add(8)
	link.travel()
	link.swap(2, 0)
	link.insert(29, 9)
	link.travel()
	link.sort()
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
		if cur.node == nil {
			fmt.Println(cur.val)
			return
		} else {
			fmt.Print(cur.val, "->")
			cur = cur.node
		}

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

func (L *Link) remove(val interface{}) {
	if L.head == nil {
		return
	}
	if L.head.val == val {
		if L.length() == 1 {
			L.head = nil
		} else {
			L.head = L.head.node
		}
	}
	cur := L.head
	for cur.node != nil && cur.node.val != val {
		cur = cur.node
	}
	if cur.node == nil {
		return
	}
	cur.node = cur.node.node

}

func (L *Link) search(val interface{}) int {
	//找到返回对应的第一个index,否则返回 -1
	// if L.head == nil {
	// 	return -1
	// }
	cur := L.head
	for count := 0; cur != nil; count++ {
		if cur.val == val {
			return count
		}
		cur = cur.node
	}
	return -1
}

func (L *Link) add(val interface{}) {
	//L.insert(val 0)
	new_node := newNode(val)
	new_node.node = L.head
	L.head = new_node
}

func (L *Link) index(index int) interface{} {
	if L.head == nil || index < 0 || index > L.length()-1 {
		// panic("out of index")
		return nil
	}
	cur := L.head
	for count := 0; cur != nil; count++ {
		if count == index {
			return cur.val
		}
		cur = cur.node
	}
	return nil
}

func (L *Link) swap(x, y int) {
	if L.length() <= 1 {
		return
	}
	if x == y {
		return
	}

	if L.length() == 2 {
		L.reverse()
	} else {
		if x > y {
			x, y = y, x  //让小的先完成替换，后面大的替换后就直接return
		}
		x_val := L.index(x)
		// y_val := L.index(y)
		cur := L.head
		// for count := 0; cur != nil; count++ {
		for count := 0; count <= L.length()-1; count++ {
			// if y > x {
			// 	if count == x {
			// 		cur.val = y_val
			// 	}
			// 	if count == y {
			// 		cur.val = x_val
			// 		return
			// 	}
			// 	cur = cur.node
			// } else {
			// 	if count == y {
			// 		cur.val = x_val
			// 	}
			// 	if count == x {
			// 		cur.val = y_val
			// 		return
			// 	}
			// 	cur = cur.node
			// }
			if count == x {
				cur.val = L.index(y)
			}
			if count == y {
				cur.val = x_val
				return
			}
			cur = cur.node

		}
	}
}

func (L *Link) sort() {
	if L.length() <= 1 {
		return
	}
	for i := 0; i < L.length()-1; i++ {
		for j := 0; j < L.length()-1-i; j++ {
			if L.index(j).(int) > L.index(j+1).(int) { //先这么写着
				L.swap(j, j+1)
			}
		}
	}
}

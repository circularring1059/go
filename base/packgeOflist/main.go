package main

import (
	"container/list"
	"fmt"
)

func main() {
	l1 := list.New()
	fmt.Println(l1)

	l1.PushBack("first") //append
	l1.PushBack("second")
	element := l1.PushBack("three")
	fmt.Println("*", element)
	l1.InsertBefore("world", element)

	l1.PushBack("end")
	element1 := l1.PushFront(1) //头插法  element  相当于index?
	element2 := l1.InsertAfter("hello", element1)
	l1.Remove(element2) //remove

	for i := l1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)

	}
}

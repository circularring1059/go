package main

import (
	"fmt"
)

//接口组合//方法组合

type write interface {
	write()
}

//空接口，任意类型
type emptyInterface interface{}

type send interface {
	send()
}

type writeSend interface {
	send
	write
}

type animal struct{}

func (a *animal) write() {
	fmt.Println("I can wrte")
}

func (a *animal) send() {
	fmt.Println("I can send")
}

// func showSlice(e interface{}) {
// 	fmt.Printf("%#v %T\n", e, e)
// 	for index, vaule := range e {
// 		fmt.Println(index, vaule)
// 	}
// }

func main() {
	var writeSendIns writeSend
	animal := &animal{}
	writeSendIns = animal
	writeSendIns.send()
	writeSendIns.write()

	//空接口使用
	// selice1 := [...]int{1, 2, 3, 4}
	// showSlice(selice1)
	var map1 = make(map[int]interface{}, 3)
	fmt.Println(len(map1), map1)
	map1[1] = 1
	map1[2] = "two"
	map1[3] = [...]int{1, 2, 3}
	fmt.Println(len(map1))
}

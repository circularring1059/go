package main

import (
	"fmt"
)

type listener interface { //后面一般加个er
	listen()
}

type watcher interface {
	watch()
}

type mp4 struct {
	brand string
}

func (m *mp4) listen() {
	fmt.Println("map4 has a function of listen\n")
}

func (m *mp4) watch() {
	fmt.Println("map4 has a function of watch\n")
}

func main() {
	fmt.Println("一个类型实现多个接口")

	var m listener
	d := &mp4{
		brand: "niuda",
	}

	m = d
	fmt.Println(m)
	m.listen()

	var m1 watcher

	m1 = d
	fmt.Println(m1)
	m1.watch()

	// var s = new(int)
	// s1 := 10
	// s = &s1
	// fmt.Println(s)

	var m2 = new(mp4)
	m2 = d
	fmt.Println(m2)
	m2.listen()
}

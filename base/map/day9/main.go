package main

import "fmt"

func main() {
	fmt.Println("start job")
	map1 := map[int]int{
		1: 1,
		2: 2,
	}
	fmt.Println(map1)

	changeMap(map1)
	fmt.Println(map1)
}

func changeMap(m interface{}) { //引用传递
	m.(map[int]int)[1] = 11
	// fmt.Println(m)
}

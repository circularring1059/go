package main

import "fmt"

func main() {
	map1 := make(map[int]int, 2)
	map1[1] = 1
	fmt.Println(len(map1))
	func(maps map[int]int) {
		maps[2] = 2
		maps[1] = 2
	}(map1)
	fmt.Println(map1)
}

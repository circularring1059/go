package main

import "fmt"

// func updateMap(map1, map[int]int) {
// 	map1[1] = 111
// 	fmt.Println(map1, map1)
// }

func main() {
	var map1 = map[int]int{1: 1}
	// updateMap(map1)
	fmt.Println(map1)

	map2 := map[int]int{1: 1}
	fmt.Println(map2)
}

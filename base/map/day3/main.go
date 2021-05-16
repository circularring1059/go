package main

import "fmt"

var map1 map[int]int
map1 = map[int]int{1:1,}
func updateMap(myMap, map[int]int) {
	myMap[1] = 111
	fmt.Println(myMap, map1)
}

func main() {
	updateMap(map1)
	fmt.Println(map1)
}

package main

import "fmt"

func main() {
	fmt.Println("Don't stop")

	//slice 中包含map
	var slice1 = make([]map[int]int, 9, 9)
	slice1[0] = make(map[int]int, 1)
	slice1[0][0] = 0
	fmt.Println(slice1)
	slice1[0][1] = 1
	fmt.Println(slice1)

	slice1[1] = map[int]int{11: 11, 22: 22}
	fmt.Println(slice1)

	//map 中包含切片
	var map1 = make(map[int][]int, 1)
	map1[0] = []int{1, 2, 3}
	fmt.Println(map1)

	map1[2] = make([]int, 2, 3)
	fmt.Println(map1)
	map1[2][0] = 3
	fmt.Println(map1)
}

/*Don't stop
[map[0:0] map[] map[] map[] map[] map[] map[] map[] map[]]
[map[0:0 1:1] map[] map[] map[] map[] map[] map[] map[] map[]]
[map[0:0 1:1] map[11:11 22:22] map[] map[] map[] map[] map[] map[] map[]]
map[0:[1 2 3]]
map[0:[1 2 3] 2:[0 0]]
map[0:[1 2 3] 2:[3 0]]*/

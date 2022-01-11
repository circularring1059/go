package main

import "fmt"

func changeMap(maps map[int]int) {
	maps[2] = 5
}

func main() {
	map1 := make(map[int]int, 2)
	map1[1] = 1
	fmt.Println(len(map1))
	func(maps map[int]int) {
		maps[2] = 2
		maps[1] = 2
	}(map1)
	fmt.Println(map1)
	changeMap(map1)
	fmt.Println(map1) //map 是引用类型

	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1) // item is only a copy of the slice element.  itme 是值拷贝
		item[1] = 2                 // This 'item' will be lost on the next iteration.
	}
	fmt.Printf("Version B: Value of items: %v\n", items2) //Version B: Value of items: [map[] map[] map[] map[] map[]]

}

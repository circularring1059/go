package main

import "fmt"

func main() {
	map1 := map[string]int{}
	map2 := map[string]int{}
	fmt.Printf("%p, %p\n", map1, map2) //0xc00007a3c0, 0xc00007a3f0

	// map3 := map2
	// fmt.Printf("%p, %p\n", map2, map3) //0xc00007a3f0, 0xc00007a3f0

	// map2 = map1
	// fmt.Printf("%p, %p, %p\n", map1, map2, map3)
	// //map2["one"] = 1
	// map2["one"] = 1               //map1 map2 has the  same address
	// fmt.Println(map1, map2, map3) //map[one:1] map[one:1] map[]
	// map3["one"] = 11
	// fmt.Println(map1, map2, map3) //map[one:1] map[one:1] map[one:11]

	map2 = map1
	map3 := map2
	fmt.Printf("%p, %p, %p\n", map1, map2, map3)
	map3["one"] = 1
	fmt.Println(map1, map2, map3) //map1、2、3 share the same address

}

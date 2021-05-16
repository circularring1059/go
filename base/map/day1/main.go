package main

import "fmt"

func main() {
	fmt.Println("Learning golang map")
	var map1 map[int]int
	fmt.Println(map1 == nil) //true
	var map2 map[string]int
	fmt.Println(map2 == nil) //trye

	map1 = map[int]int{1: 1, 2: 2}
	fmt.Printf("%v-%v\n", len(map1), map1) //2-map[1:1 2:2]

	//map2["zero"] = 0 //panic: assignment to entry in nil map
	map2 = make(map[string]int, 1) //容量为5 容量不够时会动态扩容
	fmt.Println(map2)
	map2["one"] = 1
	fmt.Println(map2) //map[one:1]
	map2["two"] = 2
	fmt.Println(map2)
	map2["two"] = 22  //修改value 值
	fmt.Println(map2) //map[one:1 two:22]

	//遍历map
	for k, v := range map2 {
		fmt.Printf("%v-%v\n", k, v)
	}

	for v := range map2 { //只接收key
		fmt.Printf("%v\n", v)
	}

	for _, v := range map2 {
		fmt.Printf("%v\n", v)
	}

}

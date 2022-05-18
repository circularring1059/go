package main

import "fmt"

func main() {
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{}
	fmt.Println(slice1, slice2)                        //[] []
	fmt.Printf("%p, %p, %p\n", slice1, slice2, slice3) //0x0, 0xcf6d60, 0xcf6d60  一样
	slice1 = append(slice1, []int{4, 5}...)            //append can not init
	fmt.Println(slice1)

	map1 := make(map[int]int, 0)
	map2 := make(map[int]int, 0)
	fmt.Printf("%p, %p", map1, map2) //0xc000076480, 0xc0000764b0 不一样

}

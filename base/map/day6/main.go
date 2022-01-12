package main

import (
	"fmt"
	"sort"
)

//map  sort
func main() {
	map1 := map[int]int{3: 1, 5: 2, 1: 4}
	fmt.Println(map1)
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	keys := make([]int, len(map1))

	i := 0
	for k, _ := range map1 {
		keys[i] = k
		i++
	}

	fmt.Println(keys)
	sort.Ints(keys)
	fmt.Println(keys)

	for _, i := range keys {
		fmt.Println(i, map1[i])
	}
}

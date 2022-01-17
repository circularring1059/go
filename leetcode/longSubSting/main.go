package main

import (
	"fmt"
)

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
func lengthOfSubString(str string) (n int) {
	map1 := make(map[string]int)
	start, max_len := 0, 0
	for k, v := range str {
		if _, ok := map1[string(v)]; !ok {
			map1[string(v)] = k

		} else {
			max_len = max(max_len, k-start)
			if map1[string(v)] >= start {
				start = map1[string(v)] + 1
				map1[string(v)] = k
			} else {
				map1[string(v)] = k
			}
		}
	}
	max_len = max(max_len, len(str)-start)
	return max_len
}

func lengthOfSubString1(str string) (n int) {
	map1 := make(map[rune]int)
	start, max_len := 0, 0
	for k, v := range str {
		if _, ok := map1[v]; !ok {
			map1[v] = k

		} else {
			max_len = max(max_len, k-start)
			if map1[v] >= start {
				start = map1[v] + 1
				map1[v] = k
			} else {
				map1[v] = k
			}
		}
	}
	max_len = max(max_len, len(str)-start)
	return max_len
}
func main() {
	var str string = "rringh"
	fmt.Println(lengthOfSubString(str))
	fmt.Println(lengthOfSubString1(str))
}

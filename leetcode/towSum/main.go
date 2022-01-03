package main

import "fmt"

func towSum1(num []int, target int) []int {
	a := make([]int, 2)
	flag := false
	s := &flag
	for i := 0; i < len(num)-1; i++ {
		if flag {
			break
		}
		fmt.Println(i)
		for j := i + 1; j < len(num); j++ {
			if num[i]+num[j] == target {
				a[0], a[1] = i, j
				*s = true
			}
		}
	}
	return a
}

func main() {
	num := []int{1, 2, 3, 4, 5}
	fmt.Println(num)
	ret := towSum1(num, 3)
	fmt.Println(ret)
}

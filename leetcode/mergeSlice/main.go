package main

import (
	"fmt"
)

func main() {
	fmt.Println("start job")
	fmt.Println(mergeSlice([][]int{
		[]int{3, 9},
		[]int{4, 9},
		[]int{0, 2},
	}))
}

func mergeSlice(s [][]int) [][]int {
	if len(s) <= 1 {
		return s
	}
	ret := make([][]int, 0)
	sort_slice(s)
	start, end := s[0][0], s[0][1]
	for i := 1; i < len(s)-1; i++ {
		left_point, right_point := s[i][0], s[i][1]
		if end > left_point {
			end = max(end, right_point)
		} else {
			ret = append(ret, []int{start, end})
			start, end = left_point, right_point
		}
	}
	ret = append(ret, []int{start, end})
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func sort_slice(s [][]int) {
	for i := 0; i < len(s)-1; i++ {
		flag := true
		for j := 0; j < len(s)-i-1; j++ {
			if s[j][0] > s[j+1][0] {
				s[j], s[j+1] = s[j+1], s[j]
				flag = false
			}
		}
		if flag {
			return
		}
	}
}

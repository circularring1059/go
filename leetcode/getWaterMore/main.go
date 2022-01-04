package main

import "fmt"

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getWaterMore(num []int) int {
	left, right := 0, len(num)-1
	maxWater := Min(num[left], num[right]) * (right - left)
	// fmt.Println(left, right)
	for left < right {
		if num[left] < num[right] {
			left++
			max_tmp := Min(num[left], num[right]) * (right - left)
			maxWater = Max(maxWater, max_tmp)
		} else {
			right--
			max_tmp := Min(num[left], num[right]) * (right - left)
			maxWater = Max(maxWater, max_tmp)
		}
	}
	return maxWater
}
func main() {
	fmt.Println(getWaterMore([]int{1, 3, 5, 5, 9, 2, 3}))
	fmt.Println(getWaterMore([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

package main

import (
	"fmt"
)

func canJump(num []int) bool {
	if len(num) <= 1 {
		return true
	}
	max := num[0]
	if num[0] <= 0 {
		return false
	}
	if num[0] >= len(num)-1 || len(num) == 1 {
		return true
	}
	for i := 1; i < len(num); i++ {
		if max < i {
			return false
		}
		if i+num[i] >= len(num) {
			return true
		} else {
			if i+num[i] > max {
				max = i + num[i]
			}
		}
	}
	return false
}

func main() {
	num := []int{2, 0, 2, 0, 7}
	num1 := []int{}
	fmt.Println(canJump(num))
	fmt.Println(canJump(num1))
	fmt.Println(Solution(num1))
	fmt.Println(Solution(num))
}

func Solution(s []int) bool {
	if len(s) <= 1 {
		return true
	}

	if s[0] == 0 {
		return false
	}
	max_jump := s[0]

	for index, val := range s {
		if max_jump < index {
			return false
		}
		max_jump = max(max_jump, index+val)
		if max_jump >= len(s)-1 {
			return true
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

package main

import "fmt"

func canJump(num []int) bool {
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
	num := []int{2, 0, 0, 0, 7}
	num1 := []int{}
	fmt.Println(canJump(num))
	fmt.Println(canJump(num1))
}

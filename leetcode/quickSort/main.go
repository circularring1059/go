package main

import "fmt"

func quickSort(num []int) []int {
	if len(num) <= 1 {
		// fmt.Println(num)
		return num
	}
	left, right := []int{}, []int{}
	for i := 1; i < len(num); i++ {
		if num[i] > num[0] {
			right = append(right, num[i])
		} else {
			left = append(left, num[i])
		}
	}
	// fmt.Println(left, right)
	return append(append(quickSort(left), num[0]), quickSort(right)...)
}

func main() {
	num := []int{5, 3, 8, 2, 8, 2, 14, 0}
	num1 := []int{}
	fmt.Println(quickSort(num))
	fmt.Println(quickSort(num1))
}

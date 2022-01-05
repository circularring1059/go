package main

import "fmt"

func choiceSort(num []int) []int {
	for i := 0; i < len(num); i++ {
		middle_index := i
		for j := i + 1; j < len(num); j++ {
			if num[j] > num[middle_index] {
				middle_index = j
			}
		}
		num[i], num[middle_index] = num[middle_index], num[i]
	}
	return num
}

func main() {
	num := []int{1, 3, 4, 2, 8, 0}
	fmt.Println(choiceSort(num))
}

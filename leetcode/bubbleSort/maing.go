package main

import "fmt"

func bubbleSort(num []int) []int {
	var flag = new(bool)
	for i := 0; i < len(num)-1; i++ {
		// fmt.Println(num[i])
		if *flag {
			break
		}

		fmt.Print(*flag)
		for j := 0; j < len(num)-1; j++ {
			*flag = true
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
				*flag = false
			}
		}
	}
	return num
}

func main() {
	num := []int{10, 3, 45, 8}
	fmt.Println(bubbleSort(num))
}

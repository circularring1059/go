package main

import "fmt"

func bubbleSort(num []int) []int {
	for i := 0; i < len(num)-1; i++ {
		// fmt.Println(num[i])
		flag := true
		fmt.Println(flag)
		for j := 0; j < len(num)-1; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
				flag = false
			}
		}
		if flag {
			break
		}

	}
	return num
}

func main() {
	num := []int{10, 3, 45, 8, 0, 2, 4, 5}
	fmt.Println(bubbleSort(num))
}

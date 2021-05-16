package main

import "fmt"

func main() {
	fmt.Println("start Bubble Sort")
	var array1 = [...]int{1, 5, 2, 9, 3, 78, 0, 56, 199, 7}
	//fmt.Println(array1)
	for i := 0; i < len(array1)-1; i++ {
		for j := 0; j < len(array1)-i-1; j++ {
			if array1[j] > array1[j+1] {
				array1[j], array1[j+1] = array1[j+1], array1[j]
			}
		}
		//fmt.Println("first", array1)
	}
	fmt.Println("排序结果:", array1)
}

package main

import "fmt"

func main() {
	fmt.Println(" to do your best")

	var slice1 []int
	// make  开辟空间
	slice1 = make([]int, 3, 5)

	slice1[2] = 2
	fmt.Println(slice1)

	// slice1[3] = 3 // error
	fmt.Println(slice1)
	//append  注意接收值还是原来的切片
	slice1 = append(slice1, 4)
	fmt.Println(slice1)
}

package main

import "fmt"

//递归函数
// func factorical(i int) int {
// 	if i == 1 {
// 		return 1
// 	}
// 	return i * factorical(i-1)
// }

// func main() {
// 	fmt.Println(factorical(3))
// 	fmt.Println(factorical(7))
// }

//斐波那契
func func1(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return func1(i-1) + func1(i-2)
}

func main() {
	fmt.Println(func1(3))
	fmt.Println(func1(4))
	fmt.Println(func1(5))
}

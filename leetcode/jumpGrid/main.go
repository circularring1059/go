package main

import "fmt"

// var (
// 	left new(int),
// 	right new(int),
// 	count int,
// )
var count int

func jumpGrid(m, n int) int {
	// var (
	// 	count, left, right int
	// )
	// fmt.Println(m, n)
	if m == 0 && n == 0 {
		// fmt.Println("over")
		count++
	}
	if m > 0 {
		jumpGrid(m-1, n)
	}

	if n > 0 {
		// fmt.Println(m, n)
		jumpGrid(m, n-1)
	}
	return count
}

func main() {
	fmt.Println(jumpGrid(2, 2))
}

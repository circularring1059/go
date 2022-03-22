package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func ops(x, y int, ops func(int, int) int) {
	fmt.Println(ops(x, y))
}

func main() {
	ops(2, 8, add)
}

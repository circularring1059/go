package main

import "fmt"

//函数作为参数
func add(x, y int) int {
	return x + y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	res := calc(1, 2, add)
	fmt.Println(res) //3
}

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

//函数作为参数
func ops(x, y int, ops func(int, int) int) {
	fmt.Println(ops(x, y))
}

//函数作为返回值
func ops1(x, y int) func(int, int) int {
	fmt.Println(x, y)
	fmt.Println("return function")
	return add
}

func main() {
	ops(2, 8, add)
	funcIns := ops1(1, 9)
	fmt.Println(funcIns(1, 3))

	//匿名函数
	sub := func(x, y int) int {
		return x - y
	}
	fmt.Println(sub(7, 2))

	func(x, y int) {
		fmt.Println(x * y)
	}(2, 8)

}

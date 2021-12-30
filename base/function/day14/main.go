package main

import (
	"fmt"
)

// 012
// 012
// 2
// 1
// 0
// 3
// 3
// 3

func f1(x int) func() {
	return func() {
		x++
		fmt.Println("x=", x)
	}
}
func main() {
	//闭包
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
		fmt.Print(i)
		if i == 2 {
			fmt.Println()
		}
	}
	//非闭包
	for i := 0; i < 3; i++ {
		defer func(b int) {
			fmt.Println(b)
		}(i)
		fmt.Print(i)
		if i == 2 {
			fmt.Println()
		}
	}

	x := f1(10)
	x() //11
	x() //12
}

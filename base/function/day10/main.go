//多闭包返回
package main

import "fmt"

func calc(x, y int) (func() int, func() int) {
	fmt.Println(x, y)
	add := func() int {
		return x + y
	}

	sub := func() int {
		return x - y
	}
	return add, sub
}

func main() {
	add, sub := calc(3, 1)
	fmt.Println(add(), sub())
}

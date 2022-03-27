package main

import (
	"fmt"
)

func show(num int) int {

	if num <= 0 {
		return 0
	}

	return show(num/10) + num%10
}

func main() {
	var num = 123
	fmt.Println(num)

	fmt.Println(show(num))
}

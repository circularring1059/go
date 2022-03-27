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

func show1(num int)int{
	var sum int
	for num > 0 {
		sum += num % 10
		num = num / 10
	}
	return sum
}

func main() {
	var num = 123
	fmt.Println(num)

	fmt.Println(show(num))

	var sum int
	for num > 0 {
		sum += num % 10
		num = num / 10
	}

	fmt.Println(sum)

	fmt.Println(show1(123))
}

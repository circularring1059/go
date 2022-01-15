package main

import (
	"fmt"
	"strconv"
)

func oneTwo(dig int) []int {
	digArry := make([]int, 0)
	s := func(a int) (c, d int) {
		return a / 10, a % 10

	}
	e, b := dig, 0
	for {
		e, b = s(e)
		digArry = append(digArry, b)
		if e == 0 {
			break
		}
	}
	return digArry
}
func main() {
	fmt.Println("one two three")
	fmt.Println(oneTwo(123))
	fmt.Println([]byte(strconv.Itoa(123)))
}

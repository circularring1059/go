package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Printf("%p\n", &a)

	b := new(int)
	fmt.Println(b)
	b = &a
	fmt.Printf("%p---%d\n", &b, *b)
	fmt.Println(b)

	c := new(int)
	fmt.Println(c)
	fmt.Println(*c)

	var a1 int
	fmt.Println(a1)

	a11 := &a1
	fmt.Println("a11", a11)

	var b1 *int
	fmt.Println(b1)
}

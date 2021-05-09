package main

import "fmt"

func main() {
	a := 1
	a++            // ++ 是单独的语句，不是运算符
	fmt.Println(a) //2
	a += 1
	fmt.Println(a) //3
	a = a + 1
	fmt.Println(a)
}

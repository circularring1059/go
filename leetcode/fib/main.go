package main

import "fmt"

func fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var a int = 4
	fmt.Println(fib(a))
}

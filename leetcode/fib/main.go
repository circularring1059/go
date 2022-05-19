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

var count = 0

func fib1(n int) int {
	if n == 0 {
		count++
		// return count
	}
	if n > 0 {
		fib1(n - 1)
	}

	if n > 1 {
		fib1(n - 2)
	}
	return count

}

func main() {
	var a int = 4
	fmt.Println(fib(a))
	fmt.Println(fib1(4))
}

package main

import "fmt"

func f1() {
	fmt.Println("exec before")
}

func f2(a int, b int) {
	fmt.Println(a, b)
	fmt.Println("old func")
}

func f3(f func(int, int), a, b int) func() {
	fmt.Println("before")
	return func() {
		f(a, b)
	}
}

func main() {
	fmt.Println("helllo go")
	f := f3(f2, 1, 2)
	f()
}

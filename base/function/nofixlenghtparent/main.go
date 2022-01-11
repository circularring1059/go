package main

import "fmt"

func test(a ...int) {
	fmt.Println(a)
}

func test1(b ...string) {
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
func main() {
	test(1, 2, 3, 4, 5, 6, 7) //[1 2 3 4 5 6 7]
	test1("stringstring", "hello")
}

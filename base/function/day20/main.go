package main

import "fmt"

func main() {
	fmt.Println("hello")
	test1()
}

func test(a *int) {
	*a = *a + 1
}

func test1() {
	b := 0
	test(&b)  ///需要修改b 就需要传指正
	fmt.Println(b)
}

package main

import (
	"fmt"
)

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x  //5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5   //6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x  //5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5  //5
}

func f5() {
	defer func() {
		fmt.Println("hello")
	}()

	defer func() {
		fmt.Println("word")
	}()
}
func main() {
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5
	f5()              //word  hello
}

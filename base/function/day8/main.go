package main

import "fmt"

//闭包
func a(i int) func() int {
	// i := 0
	// var i int
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func main() {
	c := a(2)
	c()
	c()
	c()
	c()

	fmt.Println("****")
	
	d := a(3)
	d()
	d()
	d()

	a(2) //不会输出i
}

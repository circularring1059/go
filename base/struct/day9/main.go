package main

import "fmt"

//扩展方法，只能是自定义类型

type myInt int

func (m myInt) getInformation() {
	fmt.Println("the method custom")
}

func main() {
	var a myInt = 10
	fmt.Printf("%T\n", a)
	a.getInformation()
}

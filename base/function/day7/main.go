package main

import (
	"fmt"
)

//匿名函数

func main() {
	funcAdd := func(x, y int) int {
		return x + y
	}
	fmt.Println(funcAdd(1, 2))  //3
	fmt.Printf("%T\n", funcAdd) //函数类型 func(int, int) int

	//Golang匿名函数可赋值给变量，做为结构字段，或者在 channel 里传送。

	//定义一个slice, 元素类型为func
	slice1 := [](func(x, y int) int){
		func(x, y int) int { return x + y },
		func(x, y int) int { return x - y },
	}
	fmt.Println(slice1[0](9, 9)) //18
	fmt.Println(slice1[1](4, 3))

}

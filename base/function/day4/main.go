package main

import "fmt"

//定义函数类型

type calculatioin func(int, int) int //定义计算类型

func add(x, y int) int { // calculation 类型函数
	return x + y
}

func sub(x, y int) int { //计算类型函数
	if x > y {
		return x - y
	} else {
		return y - x
	}
}

func main() {
	var add1 calculatioin
	add1 = add
	fmt.Println("Learning function")
	fmt.Printf("%T\n", add1) //main.calculatioin
	fmt.Println(add1(2, 4))  //6
}

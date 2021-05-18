package main

import "fmt"

//全局变量
var str0 string = "global variable"

func func1() {
	//局部变量
	var number1 int8 = 10
	fmt.Println(number1)
	for v, k := range str0 {
		fmt.Printf("%v %c\n", v, k)
	}
}
func main() {
	fmt.Println(str0)
	func1()
	//fmt.Println(number1)   //无法使用局部变量

	//定义一个局部与全局一样，优先使用局部变量
	str0 := "local variable"
	fmt.Println(str0)

}

package main

import "fmt"

var name string = "alibaba.com"

//无参数
func myfunc() {
	var age int = 88

	fmt.Println(name, age)

}

//形参
func myfunc1(a string, b int) {
	fmt.Println(a, b)

}

//return
func myfunc2(a string, b int) (c string, d int) {
	c = a
	d = b
	//var f string = "this is myfunc2"

	return
}
func main() {
	myfunc()
	myfunc1("who are you", 99)
}

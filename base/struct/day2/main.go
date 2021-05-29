package main

import "fmt"

//定义一个类型
type myInt int

//类型别名 rune == int32
/*type byte = uint8
type rune = int32*/

type anotherString = string

func main() {
	var a myInt
	a = 10
	fmt.Printf("a=%v type of %T\n", a, a) //10 main.myInt

	var b anotherString
	b = "golang"
	fmt.Printf("%v %T\n", b, b) //golang string
}

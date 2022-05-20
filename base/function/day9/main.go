package main

import "fmt"

//闭包复制原对象指针
// func func1() func() {
// 	var a int = 100
// 	fmt.Println(a)
// 	fmt.Printf("%p\n", &a) //0xc000014088
// 	return func() {
// 		fmt.Println(a)
// 		fmt.Printf("%p\n", &a) //0xc000014088
// 	}
// }

// func main() {
// 	funcInstance := func1()
// 	funcInstance()

// }

func func1(a int) func() {
	fmt.Println(a)
	return func() {
		a++
		fmt.Println(a)
	}
}

func main() {
	funcInstance := func1(1)
	funcInstance()
	funcInstance()
}

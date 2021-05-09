package main

import "fmt"

func main() {
	//写法一
	var number1 int = 10
	if number1 == 10 {
		fmt.Println("number1=10")
	} else if number1 == 20 {
		fmt.Println("number1=120")
	} else if number1 == 30 {
		fmt.Println("number1=30")
	} else {
		fmt.Println("number1=?")
	}
	fmt.Println("number1:", number1)

	//写法二
	if number2 := 20; number2 == 10 {
		fmt.Println("number2=10")
	} else if number2 == 20 {
		fmt.Println("number2=20")
	} else if number2 == 30 {
		fmt.Println("number2=30")
	} else {
		fmt.Println("number2=?")
	}
	//fmt.Println("number2:", number2)  number2 只能判断中使用

	//for 循环
	var number3 int = 0
	for ; number3 < 10; number3++ {
		fmt.Print(number3) //0123456789
	}
	fmt.Println()
	fmt.Println(number3)

	for number4 := 0; number4 < 10; number4++ {
		fmt.Print(number4)
		fmt.Println()
	}
	// 同样number4 不能子啊该for循环外面打印

	number5 := 0
	for number5 < 10 {
		//fmt.Println("number5")
		fmt.Print(number5)
		number5++
	}
}

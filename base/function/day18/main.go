package main

import (
	"fmt"
)

//闭包

func adder() func(x int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func adder1(x int) func(x int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	adderIns := adder()
	fmt.Println(adderIns(9)) //9
	fmt.Println(adderIns(9)) //18

	adder1Ins := adder1(9)
	fmt.Println(adder1Ins(9)) //18
	fmt.Println(adder1Ins(9)) //27

	adder2Ins := adder1(9)
	fmt.Println(adder2Ins(9)) //18
	fmt.Println(adder2Ins(9)) //27
}

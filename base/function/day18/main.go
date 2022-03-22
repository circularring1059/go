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

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
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

	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7

}

package main

import (
	"fmt"
)

func Max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
func maxLenghtValidBrackets(str string) int {
	strArry := make([]rune, 0)
	// fmt.Println(len(str))
	var ret int
	var point int
	for i, v := range str {
		if v == ')' {
			if len(strArry) == 0 {
				ret = Max(i-point, ret)
				point = i
				fmt.Println("*", ret)
			} else {
				strArry = append(strArry[0:len(strArry)-1], strArry[len(strArry)-1:len(strArry)-1]...)
			}
		} else {
			strArry = append(strArry, v)
		}
	}
	// fmt.Println(point)
	if point == 0 {
		return Max(ret, len(str)-point-len(strArry))
	} else {
		return Max(ret, len(str)-point-1-len(strArry))
	}

}

func main() {
	fmt.Println(maxLenghtValidBrackets("(()))))))))))()(()))()()()()()()"))
	// str := "ring"
	// arry := []rune(str)
	// fmt.Println(arry)
	// a1 := 1
	// a := make([]int, 0)
	// b := make([]int, 0)
	// var c = new(int)
	// fmt.Printf("%v\n%p\n%v\n", a, &b, c)
	// fmt.Printf("%+v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&a)))
	// fmt.Printf("%+v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&c)))
	// fmt.Printf("%+v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&b)))
	// fmt.Printf("%d\n", &a1)
}

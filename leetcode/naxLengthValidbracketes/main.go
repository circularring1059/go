package main

import "fmt"

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
				fmt.Println(ret)
			} else {
				strArry = append(strArry[0:len(strArry)-1], strArry[len(strArry)-1:len(strArry)-1]...)
			}
		} else {
			strArry = append(strArry, v)
		}
	}
	// fmt.Println(point)
	if point == 0 {
		return Max(ret, len(str)-point)
	} else {
		return Max(ret, len(str)-point-1)
	}

}

func main() {
	fmt.Println(maxLenghtValidBrackets("(()))))()()()()()())()"))
}

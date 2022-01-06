package main

import (
	"fmt"
	"strconv"
)

func reverseString(s string) string {
	runes := []byte(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func isPalindrome(num int) bool {
	str := strconv.Itoa(num)
	fmt.Println(str)
	if str != reverseString(str) || num < 0 {
		return false
	}
	return true
}

func isPalindrome1(num int) bool {
	fmt.Printf("%T", num)
	if num < 0 {
		return false
	}
	num1 := num
	reverse, tmp := 0, 0
	for num > 0 {
		num, tmp = test(num)
		reverse = reverse*10 + tmp
	}
	return reverse == num1
}
func test(num int) (x, y int) {
	return num / 10, num % 10
}

func main() {
	num := -121
	fmt.Println(isPalindrome(num))
	fmt.Println(test(101))
	fmt.Println(isPalindrome1(01))
}

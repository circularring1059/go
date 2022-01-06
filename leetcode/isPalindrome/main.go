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

func main() {
	num := -121
	fmt.Println(isPalindrome(num))
}

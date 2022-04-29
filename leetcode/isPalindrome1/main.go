package main

import "fmt"

func main() {
	fmt.Println("start job")
	fmt.Println(isPalindrome("ringgnir"))
	fmt.Println(isPalindrome1("ring"))
}

func isPalindrome(str string) bool {
	str_arry := []byte(str)
	if len(str_arry) <= 1 {
		return true
	}
	if str_arry[0] != str_arry[len(str_arry)-1] {
		return false
	}
	return isPalindrome(string(str_arry)[1 : len(str_arry)-1])
}

func isPalindrome1(str string) bool {
	if len(str) <= 1 {
		return true
	}

	left, right := 0, len(str)-1
	for left < right {
		if []byte(str)[left] != []byte(str)[right] {
			return false
		}
		left++
		right--
	}
	return true
}

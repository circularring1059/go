package main

import "fmt"

func main() {
	fmt.Println("start job")
	fmt.Println(isPalindrome("ringgnir"))
}

func isPalindrome(str string) bool {
	str_arry := []byte(str)
	if len(str_arry) <= 1 {
		return true
	}
    if str_arry[0] != str_arry[len(str_arry)-1] {
		return false
	}
	return isPalindrome(string(str_arry)[1: len(str_arry)-1])
}	

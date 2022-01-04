package main

import "fmt"

func validBracket(str string) bool {
	// strArry := []rune(str)
	map1 := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]rune, 0)
	for _, v := range str {
		if v == ')' || v == ']' || v == '}' && v != map1[v] {
			if len(stack) == 0 {
				return false
			} else {
				stack = append(stack[0:len(stack)-1], stack[len(stack)-1:len(stack)-1]...)
			}
		} else {
			stack = append(stack, v)
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	str := "({})"
	fmt.Println(validBracket(str))
}

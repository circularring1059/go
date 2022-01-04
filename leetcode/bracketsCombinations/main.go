package main

import "fmt"

var str string
var ret = make([]string, 0)

func gen(n int, str string, left, right int) []string {
	if len(str) == 2*n {
		ret = append(ret, str)
	}
	if left < n {
		gen(n, str+"(", left+1, right)
	}
	if right < left {
		gen(n, str+")", left, right+1)
	}
	return ret
}

func stackCombinate(n int) []string {
	return gen(n, str, 0, 0)
}

func main() {
	fmt.Println(gen(2, str, 0, 0))
	fmt.Println(stackCombinate(3))
}

package main

import "fmt"

func main() {
	sl1 := []int{1, 2, 3}
	changeSlice(sl1)
	fmt.Println(sl1)
}

func changeSlice(s []int) {  //slice  引用传递
	// s = append(s, 0)
	s[0] = 4
	fmt.Println(s)
}

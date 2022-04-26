package main 

import "fmt"

func main() {
	sl1 := []int{1,2,3}
	changeSlice(sl1)
	fmt.Println(sl1)
}

func changeSlice(sl []int){ //值拷贝
	sl = append(sl, 0)
	fmt.Println(sl)
}
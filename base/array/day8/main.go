package main

import "fmt"

func main() {
	sl1 := []int{1, 2, 3}
	changeSlice(sl1)
	fmt.Println(sl1)

	sl2 := []int{1, 2, 3}
	changeSlice1(&sl2)
	fmt.Println("sl2 is:", sl2)

}

func changeSlice(s []int) { //slice  引用传递
	// s = append(s, 0)
	s[0] = 4
	fmt.Println(s)
}

func changeSlice1(s *[]int) { //使用append 时需要传指针才能正确修改
	*s = append(*s, 2)
	fmt.Println(*s)
}

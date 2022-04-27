package main 

import "fmt"

func main () {
	fmt.Println("start job")
	sl := []int{2,7,9}
	fmt.Println(Solution(sl))  //[9 7 2]
	fmt.Println(sl) //[9 7 2]  //不是值拷贝？

}


func Solution(s []int) []int{
	if len(s) <= 1 {
		return s
	}

	for i := 0; i < len(s); i++ {
		middle_index := i
		for j := i;  j < len(s); j ++{
			//找到最大的那个数的索引
			if s[j] > s[middle_index]{
				middle_index = j
			}
		}
		s[i], s[middle_index] = s[middle_index], s[i]
	}
	return s
}
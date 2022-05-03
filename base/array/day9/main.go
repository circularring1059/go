package main

import "fmt"

func main() {
	// fmt.Println("start job")
	// sl := make([]int, 0, 10)
	// fmt.Println(sl)
	// sl = append(sl, 1)
	// fmt.Println(sl)
	fmt.Println(Solution([]int{1, 2, 3, 4}))
	fmt.Println(Solution([]int{}))
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	fmt.Println(append(s1, s2...))
}

func Solution(s []int) [][]int {
	ret := make([][]int, 0, len(s))
	if len(s) <= 1 {
		ret = append(ret, s)
		return ret
	} else {
		ret = append(ret, s)
		for i := 0; i < len(s)-1; i++ {
			c := make([]int, len(s), len(s))
			copy(c, ret[i])
			// head := ret[i]
			// head = append(head, head[0])
			// head = append(head[:0],head[1:]... )
			// ret  = append(ret, head)
			c = append(c, c[0])
			c = append(c[:0], c[1:]...)
			ret = append(ret, c)
		}
		return ret
	}
}

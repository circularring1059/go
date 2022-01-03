package main

import "fmt"

func twoSum1(num []int, target int) (x [2]int) {
	//a := make([]int, 2)
	flag := false
	s := &flag
	for i := 0; i < len(num)-1; i++ {
		if flag {
			break
		}
		fmt.Println(i)
		for j := i + 1; j < len(num); j++ {
			if num[i]+num[j] == target {
				// a[0], a[1] = i, j
				// x[0] = i
				// x[1] = j
				x = [2]int{i, j}
				*s = true
				break
			}
		}
	}
	return x
}

func twoSum2(num []int, target int) (ret [2]int) {
	map1 := make(map[int]int, len(num))
	map1[num[0]] = 0
	for i := 1; i < len(num); i++ {
		fmt.Println(map1)
		if index, ok := map1[target-num[i]]; ok {
			fmt.Println(index, i)
			ret = [2]int{index, i}
			break
		} else {
			map1[num[i]] = i
		}
	}
	return ret
}
func main() {
	num := []int{1, 2, 3, 4, 5, 5}
	// fmt.Println(num)
	ret := twoSum1(num, 7)
	fmt.Println(ret)

	num1 := []int{1, 2, 3, 4, 5, 5}
	fmt.Println(twoSum2(num1, 9))

}

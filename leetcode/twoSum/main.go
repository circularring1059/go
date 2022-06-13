package main

import "fmt"

func main() {
	fmt.Println("study")
	var sl = []int{1, 2, 4}
	fmt.Println(twoSun(sl, 6))
}

func twoSun(nums []int, target int) [][]int {
	sl := make([][]int, 0)
	mp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		_, ok := mp[target-nums[i]]
		if ok {
			sl = append(sl, []int{mp[target-nums[i]], i})
		}
		mp[nums[i]] = i
	}
	return sl

}

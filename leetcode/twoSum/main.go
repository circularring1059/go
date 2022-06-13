package main

import "fmt"

func main() {
	fmt.Println("study")
	var sl = []int{1, 2, 4}
	fmt.Println(twoSun1(sl, 3))
}

func twoSun(nums []int, target int) [][]int {
	sl := make([][]int, 0)
	mp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		//_, ok := mp[target-nums[i]]
		if _, ok := mp[target-nums[i]]; ok {
			sl = append(sl, []int{mp[target-nums[i]], i})
		}
		mp[nums[i]] = i
	}
	return sl

}

func twoSun1(nums []int, target int) [][]int {
	sl := make([][]int, 0)
	//mp := make(map[int]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[j]+nums[i] {
				sl = append(sl, []int{i, j})
			}
		}
	}
	return sl
}

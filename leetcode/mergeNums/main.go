package main

import "fmt"

func main() {
	var sl = [][]int{[]int{2, 4}, []int{1, 6}, []int{7, 8}}
	sortNums(sl)
	fmt.Println(sl)
	fmt.Println(merge(sl))
}

func merge(nums [][]int) [][]int {
	sortNums(nums)
	sl := make([][]int, 0)
	start, end := nums[0][0], nums[0][1]
	for i := 1; i < len(nums); i++ {
		if nums[i][0] <= end {
			end = max(end, nums[i][1])
		} else {
			sl = append(sl, []int{start, end})
			start, end = nums[i][0], nums[i][1]
		}
	}
	sl = append(sl, []int{start, end})
	return sl

}

func sortNums(nums [][]int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums)-1; i++ {
		mid := i
		for j := i + 1; j < len(nums); j++ {
			if nums[mid][0] > nums[j][0] {
				mid = j
			}
		}
		nums[mid], nums[i] = nums[i], nums[mid]
	}

}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

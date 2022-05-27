package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, "*", i, "=", j*i, " ")
		}
		fmt.Println()
	}

	slice1 := []int{4, 3, 5, 3, 10, 4, 6, 1, 2}
	fmt.Println(slice1)
	// bubbling(slice1)
	// choiceSort(slice1)
	slice1 = binSort(slice1)
	fmt.Println(slice1)
	slice2 := make([]int, 2, 10)
	fmt.Println(slice2)
	fmt.Println("*", binFind(slice1, 7))

	map1 := make(map[int]int, 1)
	fmt.Println(map1)
	map1[1] = 1
	fmt.Println(map1)
}

func bubbling(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		flag := true
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}

func choiceSort(nums []int) {
	for i := 0; i <= len(nums)-1; i++ {
		mix := i
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[mix] > nums[j] {
				mix = j
			}
		}
		nums[mix], nums[i] = nums[i], nums[mix]
	}
}

func binSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	middle := nums[0]
	left := make([]int, 0, len(nums))
	right := make([]int, 0, len(nums))
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] > middle {
			right = append(right, nums[i])
		} else {
			left = append(left, nums[i])
		}
	}
	return append(append(binSort(left), middle), binSort(right)...)
}

func binFind(nums []int, number int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		fmt.Println(left, right)
		middle := (right + right) / 2
		if nums[middle] == number {
			return true
		} else if nums[middle] > number {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return false
}

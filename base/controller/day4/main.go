package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, "*", i, "=", j*i, " ")
		}
		fmt.Println()
	}

	slice1 := []int{2, 0, 0, 0, 10, 4, 6, 1, 2}
	fmt.Println(slice1)
	// bubbling(slice1)
	// choiceSort(slice1)
	slice1 = binSort(slice1)
	fmt.Println(slice1)
	slice2 := make([]int, 2, 10)
	fmt.Println(slice2)
	fmt.Println("*", binFind(slice1, 7))

	// map1 := make(map[int]int, 1)
	// fmt.Println(map1)
	// map1[1] = 1
	// fmt.Println(map1)
	// fmt.Println(canJump(slice1))
	node5 := node{5, nil}
	node4 := node{4, &node5}
	node3 := node{3, &node4}
	node2 := node{2, &node3}
	node1 := node{1, &node2}
	// node5.node = &node2
	println("hello*", isLoop(&node1))
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

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	if nums[0] == 0 {
		return false
	}

	max_jump := nums[0]
	for i := 1; i <= len(nums); i++ {
		//判断是否能够跳到该点
		if max_jump < i {
			return false
		}

		max_jump = max(i+nums[i], max_jump)
		if max_jump >= len(nums)-1 {
			return true
		}
	}
	return false
}

type node struct {
	val  int
	node *node
}

func isLoop(n *node) bool {
	if n == nil || n.node == nil || n.node.node == nil {
		return false
	}

	one_point, two_point := n, n
	for two_point.node != nil {
		one_point = one_point.node
		two_point = two_point.node.node
		if two_point == nil {
			return false
		}
		if two_point == one_point || two_point.node == one_point {
			return true
		}
	}
	return false

}

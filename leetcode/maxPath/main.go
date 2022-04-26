package main

import "fmt"

func main() {
	fmt.Println("start")
	fmt.Println(maxPath(1, 1))
	fmt.Println(jumpSteps(3))
}

func backtack(x, y int, maxPath *int) {
	if x == 0 && y == 0 {
		*maxPath++
	} else {
		if x > 0 {
			backtack(x-1, y, maxPath)
		}
		if y > 0 {
			backtack(x, y-1, maxPath)
		}
	}

}

func backtack1(x, y int, number *int) {
	if x == y {
		*number++
	} else {
		if y-x >= 1 {
			backtack1(x+1, y, number)
		}
		if y-x >= 2 {
			backtack1(x+2, y, number)
		}
	}

}

func maxPath(x, y int) int {
	maxPath := 0
	backtack(x, y, &maxPath)
	return maxPath

}

func jumpSteps(m int) int {
	number := 0
	backtack1(number, m, &number)
	return number
}

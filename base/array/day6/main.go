package main

import "fmt"

func main() {
	array := [4]int{}
	slice := array[:]
	fmt.Println(array, slice) //[0 0 0 0] [0 0 0 0]

	func(array [4]int) {
		array[1] = 8
	}(array)
	fmt.Println(array, slice) //[0 0 0 0] [0 0 0 0]

	func(slice []int) {
		slice[0] = 9
	}(slice)
	fmt.Println(array, slice) //[9 0 0 0] [9 0 0 0]

}

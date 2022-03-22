package main

import "fmt"

func main() {
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{}
	fmt.Println(slice1, slice2)                        //[] []
	fmt.Printf("%p, %p, %p\n", slice1, slice2, slice3) //0x0, 0xcf6d60, 0xcf6d60
	slice1 = append(slice1, []int{4, 5}...)
	fmt.Println(slice1)
}

package main

import "fmt"

func main() {
	array1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array1)
	for i := range array1 {
		// defer fmt.Printf("%v-%v\n", i, j)
		defer fmt.Println(i)

	}
}

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------------")

	var j int = 0
	for ; j < 9; j++ {
		fmt.Println(j)
	}

	fmt.Println("--------------")
	for j < 15 {
		fmt.Println(j)
		j++
	}
}

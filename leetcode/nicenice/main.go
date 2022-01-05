package main

import "fmt"

func niceNice() {

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d  ", j, i, i*j)
		}
		fmt.Println()
	}
}

func main() {
	niceNice()
}

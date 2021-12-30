package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
		fmt.Print(i)
		if i == 2 {
			fmt.Println()
		}
	}
}

package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				break
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
	fmt.Println("****")

LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	fmt.Print("*************")
LABEL2:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				break LABEL2
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

}

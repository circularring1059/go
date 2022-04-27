package main

import "fmt"

func main() {
	fmt.Println("strt job")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i, "*", j, "=", i*j, " ")
		}
		fmt.Println()
	}


	fmt.Println("*************")

	arr := [...]int{1,2,3,4,5,6,7,8,9}
	for _,val := range arr{
		for val1 :=1; val1 <= val; val1++{
			fmt.Print(val, "*", val1, "=", val*val1, " ")
		}
		fmt.Println()
	}
}

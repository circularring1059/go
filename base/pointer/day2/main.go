package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning golang pointer")

	num1 := 10
	str1 := "helloWorld"

	var pointer1 *int
	pointer1 = &num1
	fmt.Println(pointer1)
	fmt.Println(*pointer1) //10
	fmt.Printf("%p\n", &pointer1)
	fmt.Printf("%p\n", pointer1)

	var pointer2 *string
	pointer2 = &str1
	fmt.Println(pointer2)
}

package main

import "fmt"

func main() {
	for i := 0; i <= 4; i++ {
		fmt.Print(i)
	}
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE

}

package main

import "fmt"

func myfunc() (int, string) {
	return 100, "ring"
}

func main() {
	number, _ := myfunc()
	fmt.Println(number)

}

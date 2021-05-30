package main

import "fmt"

//匿名结构体

type person struct {
	string //name
	int    // age
}

func main() {
	fmt.Println("keep moving")
	var xiaoming person = person{
		"xiaoming",
		12,
	}
	fmt.Println(xiaoming)
	fmt.Printf("%v %v\n", xiaoming.string, xiaoming.int)
}

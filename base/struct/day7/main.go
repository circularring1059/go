package main

import "fmt"

//方法

type person struct {
	name string
	age  int
}

func (p person) getName() { //形参用结构体开头小写字母
	fmt.Println(p.name)
}

func main() {
	person1 := person{
		"cia",
		19,
	}
	person1.getName()
}

package main

import "fmt"

type nameAge struct {
	string
	int
}

type nameAge1 struct {
	name string
	age  int
}
type person struct {
	nameAge
	id int
}

type cat struct {
	nameAge1
	id int
}

func main() {
	xiaohong := person{
		nameAge: nameAge{
			"xiaohong",
			12,
		},
		id: 2,
	}
	fmt.Println(xiaohong)
	fmt.Printf("%v %v\n", xiaohong.int, xiaohong.string)

	var tom cat = cat{
		nameAge1: nameAge1{
			// name: "tom",
			// age:  1,
			"tom",
			1,
		},
		id: 23,
	}
	fmt.Println(tom)
	fmt.Printf("%v %v", tom.name, tom.age) //使用匿名结构体时可以不用写tom.nameAge1.name

}

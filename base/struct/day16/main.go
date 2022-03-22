package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	bobo := person{"bobo", 17}
	fmt.Println(bobo.name, bobo.age)

	bobo1 := person{}
	bobo1.age = 18
	bobo1.name = "bobo"
	fmt.Println(bobo1.name, bobo1.age)

	var bobo2 person = person{}
	bobo2.age = 19
	bobo2.name = "bobo"
	fmt.Println(bobo2.name, bobo2.age)

	var bobo3 person
	bobo3.age = 20
	bobo3.name = "bobo"
	fmt.Println(bobo3.name, bobo3.age)

	//匿名struct

	var bobo4 struct {
		name string
		age  int
	}
	bobo4.age = 9
	bobo4.name = "bobo4"
	fmt.Println(bobo4)

	//指针struct
	var bobo5 = new(person)
	bobo5.age = 5
	bobo5.name = "bobo5"
	fmt.Println(bobo5, bobo5.name, bobo5.age)

	bobo6 := person{}
	fmt.Println(bobo6)

	var bobo7 person
	fmt.Println(bobo7)

	var bobo8 = person{
		age:  18,
		name: "bobo8",
	}

	var bobo9 = &person{
		age:  19,
		name: "bobo9",
	}

	fmt.Println(bobo8, bobo9)

	bobo10 := person{
		age: 19,
	}

	bobo11 := person{
		name: "bobo",
	}

	fmt.Println(bobo10, bobo11)

	bobo12 := person{
		"bobo12",
		19,
	}
	fmt.Println(bobo12)

}

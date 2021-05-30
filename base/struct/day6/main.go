package main

import "fmt"

//构造函数 new 开头

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

//指针类型
func newPersion1(name string, age int) *person {
	// var pe = new(person)
	// (*pe).name = name
	// pe.age = age
	// return pe
	// return &person{
	// 	name: name,
	// 	age:  age,
	// }
	// var pe *person = &person{
	// 	name: name,
	// 	age:  age,
	// }
	// return pe
	var pe *person
	pe = new(person)
	(*pe).name = name
	(*pe).age = age
	return pe
}

func main() {
	var xiaoming = newPerson("xiaoming", 19)
	var xiaohong = newPerson("xiaohong", 20)
	fmt.Printf("%v %v\n", xiaoming, xiaohong)

	var oer = newPersion1("oer", 27)
	fmt.Println(oer)
	fmt.Printf("%p\n", oer)
	fmt.Printf("%p\n", &oer)

	var cer *person
	var person1 = person{
		name: "person1",
		age:  19,
	}
	cer = &person1
	fmt.Println(cer)                     //&{person1 19}
	fmt.Printf("%p %p\n", cer, &person1) //0xc0000040d8 0xc0000040d8
}

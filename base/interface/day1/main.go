package main

import "fmt"

type person struct {
	name string
	age  int8
}
type cat struct {
	Name string
	Age  int8
}
type dog struct {
	Name string
	age  int8
}

type deam interface {
	getStruct()
}

func (p person) getStruct() {
	fmt.Println("this in a person dstruct")
}

func (c cat) getStruct() {
	fmt.Println("this is a cat struct")
}

func (d dog) getStruct() {
	fmt.Println("this is a dog struct")
}

func getName(x deam) {
	x.getStruct()
}
func main() {
	fmt.Println("Learning interface")
	var person1 = person{
		"persoon1",
		12,
	}

	var cat1 = cat{
		"cat1",
		1,
	}

	var dog1 = dog{
		"dog1",
		3,
	}

	getName(person1)
	fmt.Println(person1, cat1, dog1)

	getName(person1)
	getName(cat1)
	getName(dog1)
}

// package main

// import "fmt"

// type Inter interface {
// 	Get()
// }

// type St struct {
// 	Age int
// }

// func (s St) Get() {
// 	fmt.Println("helloworld")
// }

// // func (s *St) Set(age int) {
// // 	s.Age = age
// // }

// func test(i Inter) {
// 	i.Get()
// }

// func main() {
// 	s := St{12}
// 	test(s)

// }

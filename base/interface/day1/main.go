// package main

// import "fmt"

// type person struct {
// 	name string
// 	age  int8
// }
// type cat struct {
// 	name string
// 	age  int8
// }
// type dog struct {
// 	name string
// 	age  int8
// }

// type deam interface {
// 	getStrurt() string
// }

// func (p person) getStruct() {
// 	fmt.Print("this in a person dstruct")
// }

// func (c cat) getStruct() {
// 	fmt.Println("this is a cat struct")
// }

// func (d dog) getStruct() {
// 	fmt.Println("this is a dog struct")
// }

// func getName(x deam) {
// 	x.getStrurt()
// }
// func main() {
// 	fmt.Println("Learning interface")
// 	var person1 = person{
// 		"persoon1",
// 		12,
// 	}

// 	var cat1 = cat{
// 		"cat1",
// 		1,
// 	}

// 	var dog1 = dog{
// 		"dog1",
// 		3,
// 	}

// 	// getName(person1)
// 	fmt.Println(person1, cat1, dog1)
// 	getName(person1)

// }

package main

import "fmt"

type Inter interface {
	Get() int
	Set(int)
}

type St struct {
	Age int
}

func (s St) Get() int {
	return s.Age
}

func (s *St) Set(age int) {
	s.Age = age
}

func test(i Inter) {
	i.Set(10)
	fmt.Println(i.Get())
}

func main() {
	s := St{}
	test(&s)
}

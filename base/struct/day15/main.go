package main

import "fmt"

type struct1 struct {
	a int
	b string
}

func (s struct1) changeName() {
	s.b = s.b + "change" //值拷贝
	fmt.Println(s.b)
}

func (s *struct1) changeName1() {
	s.b = s.b + "change" // 引用传递
	fmt.Println(s.b)
}

var a1 struct1

func main() {

	fmt.Println(a1)
	a := struct1{1, "ring"}
	fmt.Println(a)
	a.changeName()
	fmt.Println(a) //{1 ring}
	a.changeName1()
	fmt.Print(a) //{1 ringchange}

}

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

func main() {
	a := struct1{1, "ring"}
	fmt.Println(a)
	a.changeName()
	fmt.Println(a)
	a.changeName1()
	fmt.Print(a)

}

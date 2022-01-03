package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	var a4 = new(int)
	var a3 = new(student)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Printf("%p\n", &a3)
	fmt.Printf("%p\n", a3)
	fmt.Printf("a4:%d\n", a4)
	// var c int = 15
	// var b = new(int)
	// // var a = new(student)
	// // fmt.Println(&a)
	// // fmt.Println("****")
	// // fmt.Println(&b)
	// // fmt.Printf("%x\n", &a)
	// fmt.Printf("%p\n", &b)
	// fmt.Printf("%x\n", b)
	a := 10
	b := &a
	fmt.Printf("%x\n", b)
	fmt.Printf("%p\n", b)
	// fmt.Printf("%d\n", &a)
	// fmt.Print(&c)

	// var a1 *student
	a2 := new(int)
	a1 := new(student)
	a1.name = "ring"
	a1.age = 15
	fmt.Println(a2)
	fmt.Printf("%x\n", *a1)
	a1.name = "ring"
	a1.age = 18
	// fmt.Printf("%#v", a1)
	fmt.Printf("%x", "ring")
	for k, v := range []byte("ring") {
		fmt.Println(k, v)
	}
}

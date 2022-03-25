package main

import (
	"fmt"
	"reflect"
)

type person struct {
	age  int
	name string
}

func createPerson(age int, name string) *person {
	return &person{
		age,
		name,
	}
}

func main() {
	fmt.Println("package of reflect")

	var a = 1
	var b = "b"
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))

	c := float32(3)
	fmt.Printf("%T\n", c)
	// fmt.Println(reflect.Value(c))

	d := reflect.TypeOf(c)
	fmt.Println(d, d.Name(), d.Kind())

	var f interface{} = 1.2117
	var g interface{} = []interface{}{"hello", 2, map[int]string{1: "one"}}
	fmt.Println(reflect.TypeOf(f)) // float64
	fmt.Println(reflect.TypeOf(g)) // float64

	//Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
	alice := createPerson(13, "alice")
	typeOfAlice := reflect.TypeOf(alice)
	fmt.Println(reflect.TypeOf(alice))
	fmt.Printf("Name:%v, Kind:%v\n", typeOfAlice.Name(), typeOfAlice.Kind()) //Name:, Kind:ptr

	var bobo interface{} = person{18, "bobo"}
	typeOfBobo := reflect.TypeOf(bobo)
	fmt.Println(reflect.TypeOf(bobo))                                      //main.person
	fmt.Println(typeOfBobo.Name(), typeOfBobo.Kind())                      //person struct
	fmt.Printf("Name:%v, Kind:%v\n", typeOfBobo.Name(), typeOfBobo.Kind()) //Name:person, Kind:struct
}

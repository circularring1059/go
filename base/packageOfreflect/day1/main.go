package main

import (
	"fmt"
	"reflect"
)

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
	fmt.Println(reflect.TypeOf(f))

}

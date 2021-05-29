//匿名结构体
package main

import "fmt"

func main() {
	var maria struct {
		name, gender string
		age          int
	}
	maria.age, maria.gender = 20, "女"
	maria.name = "maria"
	fmt.Println(maria) //{maria 女 20}
	maria.age = 21
	fmt.Printf("%#v", maria) //struct { name string; gender string; age int }{name:"maria", gender:"女", age:21}
}

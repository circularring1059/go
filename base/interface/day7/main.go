package main

import "fmt"

type person struct{}

type dog struct{}

type cat struct{}

type speak interface {
	//struct  需要实现interface 中*所有*方法或者其嵌套的struct 能够实现其中的方法也可
	speak()
}

func (p person) speak() {
	fmt.Println("I am  person")
}

func (d *dog) speak() {
	fmt.Println("dog")
}

func (c *cat) speak() {
	fmt.Println("cat")
}

func main() {
	//一个类型可以同时实现多个接口
	//多种类型实现同一接口
	fmt.Println("interface review")
	personIns := person{}
	personIns.speak()

	var dogIns dog
	dogIns.speak()
	var catIns = new(cat)
	catIns.speak()

	var interfaceIns speak
	personIns1 := person{} //可以传指针，也可以传struct

	interfaceIns = personIns1
	interfaceIns.speak()
	fmt.Printf("%T\n", interfaceIns)

	var dogIns1 dog
	interfaceIns = &dogIns1 //必须传指针，不能是dog struct
	interfaceIns.speak()
	fmt.Printf("%T\n", interfaceIns)

	var catIns1 = new(cat)
	interfaceIns = catIns1
	catIns1.speak()
	fmt.Println(interfaceIns) //&{}

}

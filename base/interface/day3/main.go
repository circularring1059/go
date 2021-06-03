package main

import "fmt"

type run interface {
	run(string)
}

type cat struct {
	name string
	age  int
}

func (c cat) run(arg string) {
	fmt.Printf("%v is runing\n", arg)
}

type eat interface {
	eat(string)
}

type dog struct {
	name string
	age  int8
}

//这里使用的是指针，后面赋值时就得传指针类型
func (d *dog) eat(arg string) {
	fmt.Printf("%v is eatting\n", arg)
}

func main() {
	fmt.Println("Learing interface")
	fmt.Println(">>>>>>>>>>>>>>>>>")

	cat1 := cat{
		"cat1",
		1,
	}
	fmt.Println(cat1)
	var i run
	i = cat1
	i.run(cat1.name)

	//这里是指针类型也ok
	cat2 := &cat{
		"cat2",
		2,
	}
	fmt.Println(cat2)
	fmt.Println(*cat2) //{cat2 2}
	var i2 run
	i2 = cat2
	i2.run(cat2.name) //cat2 is runing

	var d1 eat
	dog1 := dog{
		"dog1",
		1,
	}
	d1 = &dog1 //这里要取地址
	fmt.Println(d1)
	d1.eat(dog1.name)

	dog2 := &dog{
		"dog2",
		2,
	}
	var d2 eat
	d2 = dog2
	fmt.Println(d2)
	d2.eat(dog2.name)

}

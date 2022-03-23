package main

import "fmt"

//结构体中字段匿名
type person struct{
	personalId int
	int
	string
}


//结构体嵌套
type dog struct{
	sex string
	age int
}
type animal struct{
	one int
	two  dog
}
//结构体继承

func main() {
	var bobo1 person = person{110, 17, "bobo1"}
	fmt.Println(bobo1, bobo1.int, bobo1.string)

	bobo2 := animal{1, dog{"male", 4}} //{17 bobo1} 17 bobo1
	fmt.Println(bobo2, bobo2.two.sex) //{1 {male 4}} male
}
package main

import "fmt"

type person struct{}

type dog struct{}

type cat struct{}

func (p *person) speak() {
	fmt.Println("I am  person")
}

func (d *dog) speak() {
	fmt.Println("dog")
}

func (c *cat) speak() {
	fmt.Println("cat")
}

func main() {
	fmt.Println("interface review")
	personIns := person{}
	personIns.speak()

	var dogIns dog
	dogIns.speak()
	var catIns = new(cat)
	catIns.speak()
}

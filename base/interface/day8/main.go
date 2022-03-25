package main

import (
	"fmt"
)

type animal struct{}

type pig struct {
	*animal
}

type action interface {
	speak()
	eat()
}

func (a *animal) speak() {
	fmt.Println("I can speack")
}

func (p *pig) eat() {
	fmt.Println("I can eat")
}

func main() {
	pigIns := &pig{&animal{}}
	pigIns.eat()
	pigIns.speak()
	var interfaceIns action
	interfaceIns = pigIns
	interfaceIns.eat()
	interfaceIns.speak()

}

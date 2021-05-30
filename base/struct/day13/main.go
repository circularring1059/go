package main

import "fmt"

//继承

type pet struct {
	name string
}

type cat struct {
	id  int
	pet //继承
}

func (p pet) getName() {
	fmt.Printf("my name is %v\n", p.name)
}

func (c cat) getId() {
	fmt.Printf("my id is %v\n", c.id)
}

func main() {
	tom := cat{
		id: 1,
		pet: pet{
			"tom",
		},
	}
	fmt.Println(tom)
	tom.getName()
	tom.getId()
}

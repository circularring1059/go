package main

import (
	"fmt"
)

type calc interface {
	add(x, y int) int

	// sub(x, y int) (z int) {
	// 	if  x > y {
	// 		return x -y
	// 	}else {
	// 		return y -x
	// 	}
	// }

}

type computer struct {
	brand string
}

func (c computer) add(x, y int) (z int) {
	return x + y
}

func h(i calc, x int, y int) {
	i.add(x, y)
}

func main() {
	fmt.Println("Try to do it")
	var leavel = computer{
		"leavle",
	}

	var t calc
	t = leavel
	fmt.Println(t.add(3, 7))

	var t1 calc
	fmt.Printf("%T\n", t1) //<nil>
	t1 = leavel
	fmt.Printf("%T\n", t1) //main.computer
}

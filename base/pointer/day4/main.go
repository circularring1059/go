package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Printf("a:%d prt:%p \n", a, &a)
	fmt.Printf("b:%p type:%T \n", b, b)
	fmt.Println(&b)
}

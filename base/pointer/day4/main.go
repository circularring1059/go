package main

import "fmt"

func main() {
	// a:10 prt:0xc0000ac058
	// b:0xc0000ac058 type:*int
	// 0xc0000d8018
	a := 10
	b := &a
	fmt.Printf("a:%d prt:%p \n", a, &a)
	fmt.Printf("b:%p type:%T \n", b, b)
	fmt.Println(&b)
}

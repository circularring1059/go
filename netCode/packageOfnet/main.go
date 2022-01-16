package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:9201")
	if err != nil {
		fmt.Println("Listen fialed")
	}

	con, err := listen.Accept()
	if err != nil {
		fmt.Println("accept failed")

	}
	var tmp [1238]byte
	n, err := con.Read(tmp[:])
	if err != nil {
		fmt.Println("")
	}
	fmt.Println("start net work")

	fmt.Println(n)
}

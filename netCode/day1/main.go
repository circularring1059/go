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
		return
	}
	defer con.Close()
	for {
		var msg [128]byte
		n, err := con.Read(msg[:])
		if err != nil {
			fmt.Println("")
		}
		fmt.Println(string(msg[:n]))
	}

}

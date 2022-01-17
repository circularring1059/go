package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client")
	con, err := net.Dial("tcp", "127.0.0.1:9201")
	if err != nil {
		fmt.Println("connet failed")
	}
	defer con.Close()
	fmt.Println("start send msg")
	for {
		_, err = con.Write([]byte("Hello world"))
		if err != nil {
			fmt.Println("send msg failed")
		}
		time.Sleep(time.Second * 4)
	}
}

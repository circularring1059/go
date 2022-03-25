package main

import (
	"fmt"
)

//接口组合

type write interface {
	write()
}

type send interface {
	send()
}

type writeSend interface {
	send
	write
}

type animal struct{}

func (a *animal) write() {
	fmt.Println("I can wrte")
}

func (a *animal) send() {
	fmt.Println("I can send")
}

func main() {
	var writeSendIns writeSend
	animal := &animal{}
	writeSendIns = animal
	writeSendIns.send()
	writeSendIns.write()
}

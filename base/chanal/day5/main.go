package main

import (
	"fmt"
)

var chan1 chan int 

func main () {
	//需要使用make(引用类型)开辟空间
	chan1 := make(chan int ) //无缓冲
	fmt.Println(chan1)

	chan2 := make(chan string, 5) //有缓冲
	fmt.Println(chan2)

	chan1 <- 1   //chan 无缓冲，必需有接受者，否则deadlock
	fmt.Println("***")
	
}
package main

import (
	"fmt"
	"time"
)

var chan1 chan int 

func main () {
	//需要使用make(引用类型)开辟空间
	chan1 := make(chan int ) //无缓冲
	fmt.Println(chan1)

	chan2 := make(chan string, 5) //有缓冲
	fmt.Println(chan2)

	chan2  <- "chan"   //chan 无缓冲，必需有接受者，否则deadlock
	fmt.Println("***")

	var chan3 chan int
	go func() {
		fmt.Println("before")
		x := <-chan3  //阻塞
		fmt.Println(x)
		fmt.Println("after")
	}()
	// chan3 <- 2  //需要开辟空间才能使用
	
	time.Sleep(time.Second * 5)
	
}
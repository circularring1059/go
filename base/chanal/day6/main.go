package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var onece sync.Once

func send(c chan int) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println("intput chan2")
		c <- i //无缓chan, 需要等待receiver，接收一个值后才能返回 
		fmt.Println("after intput")
	}
	close(c)
}

func receiver(c chan int, c1 chan int) {
	defer wg.Done()
	time.Sleep(time.Second * 10)
	for i := range c {
		c1 <- i * i
		fmt.Println("intput chan3")
	}
	onece.Do(func (){close(c1)}) ///只关闭一次chan
}

func main() {
	wg.Add(3)
	chan1 := make(chan int)
	go func() { //接受chan 值
		defer wg.Done()
		fmt.Println(<-chan1)
	}()
	chan1 <- 1
	// chan1 <- 1 // deadlock

	chan2 := make(chan int, ) //有接收者，可以不用缓冲
	chan3 := make(chan int, 5)
	go send(chan2)
	go receiver(chan2, chan3)
	go receiver(chan2, chan3)
	
	for i := range chan2{
		fmt.Println("i is:", i)
	}
	
	wg.Wait()
}

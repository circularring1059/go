package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main () {
	wg.Add(1)
	chan1 := make(chan int)
	go func(){  //接受chan 值
		defer wg.Done()
		fmt.Println(<-chan1)
	}()
	chan1 <- 1
	// chan1 <- 1 // deadlock
	wg.Wait()
}
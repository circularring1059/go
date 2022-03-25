package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func run(i int) {
	time.Sleep(time.Second * 3)
	fmt.Println("xxx is  running")
	fmt.Println(i)
	wg.Done()
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go run(i)
		wg.Wait()
	}
}

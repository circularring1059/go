package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker() {
	time.Sleep(time.Second * 3)
	fmt.Println("***")
	wg.Done()
}

func main() {
	wg.Add(1)
	fmt.Println("start worker")
	go worker()
	fmt.Println("wait")
	wg.Wait()
	fmt.Print("end worker")

}

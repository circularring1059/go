//select

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(100)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	// Received on channel 1: 0
	// Received on channel 2: 5
	// Received on channel 1: 2
	// Received on channel 1: 4
	// Received on channel 1: 6
	// Received on channel 1: 8
	// Received on channel 1: 10
	// Received on channel 1: 12
	// Received on channel 1: 14
	// Received on channel 1: 16
	// Received on channel 1: 18
	// Received on channel 1: 20
	// Received on channel 1: 22
	// Received on channel 1: 24
	// Received on channel 1: 26
	// Received on channel 1: 28
	// Received on channel 1: 30
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}

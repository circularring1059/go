package main

import (
	"fmt"
	"time"
)

// func main() {
// 	fmt.Println("chanal")
// 	// ch := make(chan int) //无缓冲
// 	// ch <- 2              //error

// 	// a := <-ch  //error
// 	// fmt.Println(a)
// 	// fmt.Println("*")
// 	// ch <- 3
// 	// fmt.Println("**")
// }

// package main

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string
	time.Sleep(0.5 * 1e9)
	for {
		input = <-ch
		fmt.Printf("%s\n", input)
	}
}

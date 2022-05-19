package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	fmt.Println("close chanal")
	close(ch)
}

func getData(ch chan string) {
	time.Sleep(time.Second * 5)
	// for {
	// 	input, open := <-ch
	// 	if !open { //  通道关闭或者阻塞
	// 		break
	// 	}
	// 	fmt.Printf("%s\n", input)
	// }
	for i := range ch { //  for range  可以自动的检测chanal  是否关闭
		fmt.Println(i)
	}
	fmt.Println("channel is closed")
}

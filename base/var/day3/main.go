package main

import "fmt"

func myfunc() (int, string) {
	return 100, "ring"
}

func main() {
	number, _ := myfunc() // _  接收匿名变量，占位，不占内存
	fmt.Println(number)

}

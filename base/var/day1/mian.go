package main

// import "fmt"
import (
	"fmt"
)

var name string = "ring"
var age int = 24 // 全局变量可以声明不使用

func main() {
	fmt.Println(name)
	// fmt.Println(age)

	var ageIn int = 26 //局部变量声明后必须使用
	fmt.Println(ageIn)
}

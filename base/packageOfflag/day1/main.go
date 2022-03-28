package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	fmt.Println("packger flag")
	if len(os.Args) > 1 { //有个默认的  c:\Users\rice\go\src\github.com\go\base\packageOflag\day1\__debug_bin.exe
		//以空格区别每个元素
		fmt.Println(len(os.Args), os.Args, reflect.TypeOf(os.Args)) //type: []string
		for index, vaule := range os.Args {
			fmt.Println(index, vaule)
		}
	} else {
		fmt.Println("no input")
	}
}

package main

import (
	"flag"
	"fmt"
	"time"
)

var names string
var age int
var married bool
var delay time.Duration

func main() {
	fmt.Println("study package of flag")
	name := flag.String("ring", "原件", "备注")
	fmt.Printf("%v-%T\n", *name, *name) //string

	flag.StringVar(&names, "names", "原件", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间")
	fmt.Println(names)

}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := "  			ring"
	str2 := "ingstudyi"

	str3 := str1 + str2

	fmt.Println(str3)
	fmt.Println(strings.HasPrefix(str1, "rin"))      // true
	fmt.Println(strings.HasSuffix(str2, "uy"))       //false
	fmt.Println(strings.HasSuffix(str2, "y"))        //true
	fmt.Println(strings.Index(str3, "ing"))          //1
	fmt.Println(strings.LastIndex(str3, "ig"))       //4   找不到返回 -1
	fmt.Println(strings.Contains(str1, "t"))         //false
	fmt.Println(strings.Count(str3, "i"))            //3
	fmt.Println(strings.Replace(str3, "i", "I", 1))  //替换一次
	fmt.Println(strings.Replace(str3, "i", "I", -1)) //替换全部
	fmt.Println(strings.TrimSpace(str1))             //去掉开头结尾的空格或者tab

	var a = 97
	var b = 97
	fmt.Println(string(a)) //a
	c := strconv.Itoa(b)
	fmt.Printf("%T %v", c, c)
}

package main

import (
	"fmt"
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
	fmt.Println(strings.LastIndex(str3, "ing"))      //4
	fmt.Println(strings.Contains(str1, "t"))         //false
	fmt.Println(strings.Count(str3, "i"))            //3
	fmt.Println(strings.Replace(str3, "i", "I", 1))  //替换一次
	fmt.Println(strings.Replace(str3, "i", "I", -1)) //替换全部
	fmt.Println(strings.TrimSpace(str1))

}

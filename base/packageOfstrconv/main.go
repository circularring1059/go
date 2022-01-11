package main

import (
	"fmt"
	"strconv"
)

var a byte = 97

func main() {
	var str1 string = "ring"
	fmt.Print(str1, "\n")
	fmt.Println(strconv.ParseBool("1")) //true  <nil>
	fmt.Println(strconv.ParseBool(" ")) // " ": invalid syntax
	fmt.Println(strconv.ParseBool("0")) //false  <nil>
	fmt.Println(strconv.ParseBool("True"))
	fmt.Println(strconv.ParseBool("TRue")) //false strconv.ParseBool: parsing "TRue": invalid syntax

	fmt.Println("***")

	ret := make([]byte, 0)
	ret = strconv.AppendBool(ret, 0 > 1)
	fmt.Printf("%s\n", ret) //fase

	fmt.Println(strconv.FormatBool(0 < 1)) // true
	fmt.Println(strconv.FormatBool(0 > 1)) // false

	// ParseInt 将字符串转换为 int 类型
	// s：要转换的字符串
	// base：进位制（2 进制到 36 进制）
	// bitSize：指定整数类型（0:int、8:int8、16:int16、32:int32、64:int64）
	// 返回转换后的结果和转换时遇到的错误
	// 如果 base 为 0，则根据字符串的前缀判断进位制（0x:16，0:8，其它:10）
	fmt.Println(strconv.ParseInt("123", 10, 8))
	fmt.Println(strconv.ParseUint("FF", 16, 8)) //255
	fmt.Println(strconv.Atoi("2147483647"))
	fmt.Println(strconv.Quote(`C:\Windows`))
	s := "0.12345678901234567890"
	f, err := strconv.ParseFloat(s, 32)
	fmt.Println(f, err)          // 0.12345679104328156
	fmt.Println(float32(f), err) // 0.12345679
	f, err = strconv.ParseFloat(s, 64)
	fmt.Println(f, err) // 0.12345678901234568

	fmt.Println([]rune(str1))
	fmt.Printf("%c\n", a)
	fmt.Printf("%T\n", a)

}

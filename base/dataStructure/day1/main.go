/*uint8	无符号 8位整型 (0 到 255)
uint16	无符号 16位整型 (0 到 65535)
uint32	无符号 32位整型 (0 到 4294967295)
uint64	无符号 64位整型 (0 到 18446744073709551615)
int8	有符号 8位整型 (-128 到 127)
int16	有符号 16位整型 (-32768 到 32767)
int32	有符号 32位整型 (-2147483648 到 2147483647)
int64	有符号 64位整型 (-9223372036854775808 到 9223372036854775807)*/

/*uint	32位操作系统上就是uint32，64位操作系统上就是uint64
int	32位操作系统上就是int32，64位操作系统上就是int64 尽量不使用int 和uint
uintptr	无符号整型，用于存放一个指针*/

package main

import "fmt"

func main() {
	//go 语言不能表示二进制数
	//定义十进制数
	var number1 int = 10

	//使用 "_"  分割数字

	var number4 int = 1_2_3

	fmt.Println(number4)

	//定义八进制数
	var number2 int = 077

	//定义十六进制数

	var number3 int = 0x0123

	fmt.Println(number1, number2, number3)

	fmt.Printf("%b\n", number1) //十进制转二进制
	fmt.Printf("%o\n", number1) //十进制转八进制
	fmt.Printf("%x\n", number1) //十进制转十六进制

	//使用  "%T"  查看  数据类型， python(type)

	fmt.Printf("%T\n", number1)
	fmt.Printf("%T\n", number2)
	fmt.Printf("%T\n", number3)
}

/*Go语言中使用&字符放在变量前面对变量进行“取地址”操作。
Go语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：*int、*int64、*string等。*/
//"&"（取地址）和"*"（根据地址取值）
package main

import "fmt"

func main() {
	fmt.Println("Learning pointer")
	number1 := 10
	number1Addr := &number1
	fmt.Println(number1Addr)                      //0xc000014088
	fmt.Println(*number1Addr)                     //10
	fmt.Printf("%d==%d\n", &number1, number1Addr) //824633802888==824633802888 十进制
	fmt.Printf("%v==%v\n", &number1, number1Addr) //824633802888==824633802888 十六进制

	//定义一个指针类型
	var addr1 *int
	fmt.Println(addr1) //<nil>
	// *addr1 = 99
	// fmt.Println(*addr1) //panic: runtime error: invalid memory address or nil pointer dereference
	var addr2 = new(int)
	*addr2 = 123
	fmt.Println(*addr2) //123
}

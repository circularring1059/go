/*uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
rune类型，代表一个 UTF-8字符。实际是int32*/

package main

import "fmt"

//字符表示 使用''
var a = 'h'
var b = '中'
var c int32 = 20013

func main() {
	fmt.Println(a, b)           //102 20013
	fmt.Printf("%T %T\n", a, b) //int32 int32

	fmt.Printf("%c %c\n", a, b) //h 中
	fmt.Printf("%c %v\n", a, b) //h 20013
	fmt.Printf("%c\n", c)       //中  unicode 中 == 20013
	fmt.Println(c)

	// 遍历字符串
	s := "hello你好"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i]) //104(h) 101(e) 108(l) 108(l) 111(o) 228(ä) 189(½) 160( ) 229(å) 165(¥) 189(½)   一个中文占三个字节
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r) //104(h) 101(e) 108(l) 108(l) 111(o) 20320(你) 22909(好)
	}
	fmt.Println()

	//字符串修改，重新分配内存
	str1 := "helloWorld"
	str2 := "你好单词"
	byteStr1 := []byte(str1)
	byteStr1[0] = 'H'
	runeStr2 := []rune(str2)
	runeStr2[0] = '您'

	fmt.Println(string(byteStr1), string(runeStr2))
}

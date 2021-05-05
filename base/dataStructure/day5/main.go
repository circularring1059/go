/* string
\r	回车符（返回行首）
\n	换行符（直接跳到下一行的同列位置）
\t	制表符
\'	单引号
\"	双引号
\\	反斜杠
*/

package main

import "fmt"

var str1 = "hello go"

var str2 = "hello world"

func main() {

	str3 := "hello ring"
	fmt.Println(str1, str2, str3)

	//定义一个多行字符
	str4 := `one
	two
	three
	four`

	fmt.Print(str4, "\n")
	fmt.Println("--------")
	/*
	one
		two
		three
		four
	--------
	*/
}

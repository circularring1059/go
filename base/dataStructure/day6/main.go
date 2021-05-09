// string  operate
/*
方法	介绍
len(str)	求长度
+或fmt.Sprintf	拼接字符串
strings.Split	分割
strings.contains	判断是否包含
strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
strings.Index(),strings.LastIndex()	子串出现的位置
strings.Join(a[]string, sep string)	join操作
*/

package main

import (
	"fmt"
	"strings"
)

var str1 string = "taobao.com"

func main() {
	fmt.Println(str1, "lenght:", len(str1))

	str2 := "dir"
	str3 := "file"
	str4 := str2 + str3 // "+" 进行字符串拼接， 和python 一样

	fmt.Println(str4)
	fmt.Printf("%s%s", str2, str3) //dirfile

	//使用fmt.Sprintf 拼接
	str5 := fmt.Sprintf("%s%s%s", str2, str3, str4)

	fmt.Println(str5) //dirfiledirfile

	str6 := `D:\language\go\go\src` //使用反引号 "\" 就是单纯的 \  不需要要转义， 使用 ""  标识的字符串则要
	//字符串切割

	str7 := strings.Split(str6, "")  //  [D : \ l a n g u a g e \ g o \ g o \ s r c]
	str8 := strings.Split(str6, " ") //  [D:\language\go\go\src]  注意和上面的区别，和python  的splite 有点不同

	fmt.Println(str7)
	fmt.Println(str8)

	fmt.Println(strings.Split(str6, "\\")) // [D: language go go src]

	//包含 Contains  python 用 in
	fmt.Println(strings.Contains(str1, "taobao")) //true
	fmt.Println(strings.Contains(str1, "ring"))   //false
	fmt.Println(strings.Contains(str1, ""))       //true
	fmt.Println(strings.Contains(str1, " "))      //false

	//前缀、后缀
	//str1 = taobao.com
	fmt.Println(strings.HasPrefix(str1, "tao"))  //ture
	fmt.Println(strings.HasPrefix(str1, " tao")) //false  空格看成一个特殊字符处理就好
	fmt.Println(strings.HasSuffix(str1, "com"))  //true
	fmt.Println(strings.HasSuffix(str1, "com ")) //false

	//Index  和python  差不多

	//str1 = taobao.com

	fmt.Println(strings.Index(str1, "t"))     // 0
	fmt.Println(strings.Index(str1, "tao"))   // 0
	fmt.Println(strings.Index(str1, "ato"))   // -1   ？
	fmt.Println(strings.Index(str1, "o"))     // 2
	fmt.Println(strings.LastIndex(str1, "o")) // 8

	// join 和python 差不多
	// str1 = taobao.com
	fmt.Println(strings.Join(str7, "+")) // ？ D+:+\+l+a+n+g+u+a+g+e+\+g+o+\+g+o+\+s+r+c
	fmt.Println(strings.Join(str7, ""))  // ？D:\language\go\go\src

}

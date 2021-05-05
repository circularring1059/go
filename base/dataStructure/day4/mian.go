/* bool 类型 */
// true   真
// false jia

/* 布尔类型变量的默认值为false。
Go 语言中不允许将整型强制转换为布尔型.
布尔型无法参与数值运算，也无法与其他类型进行转换*/

package main

import "fmt"

func main() {
	var bool1 bool = true
	var bool2 bool // 默认false

	fmt.Println(bool1, bool2)

}

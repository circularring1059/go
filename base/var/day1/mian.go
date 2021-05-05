package main

// import "fmt"   引入多个包时可以使用下面的方式
import (
	"fmt"
)

/*声明变量方式一*/
//声明全局变量时不可以省略var关键字， 变量不能重复声明
var name string = "ring"
var age int = 24 // 全局变量可以声明不使用

func main() {
	fmt.Println(name)
	// fmt.Println(age)

	var ageIn int = 26 //局部变量声明后必须使用
	fmt.Println(ageIn)

	//方式二
	var name1 string
	var age1 int
	name1 = "ring.com"
	age1 = 27

	fmt.Println(name1)
	fmt.Println(age1)

	//方式三 不声明数据类型 类似Python
	var name2 = "ring.qq.cm"
	var age2 = 28
	fmt.Println(name2)
	fmt.Println(age2)

	//方式四 省略 var  关键字
	name3 := "163.com"
	age3 := 29

	fmt.Println(name3)
	fmt.Println(age3)

	//声明多个变量
	var age4, agee5 int = 20, 21
	fmt.Println(age4, agee5)

	//多行声明变量/
	var (
		name5 string  = "126.com"
		age6  int     = 12
		isOk  bool    = false
		dec   float32 = 1.23
	)
	fmt.Println(name5, age6, isOk, dec)
}

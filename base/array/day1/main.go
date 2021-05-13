package main

import (
	"fmt"
)

var array9 = [...]int{1, 3, 4, 5}

//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值
func printArray(myarray [4]int) {
	fmt.Println(myarray)
	myarray[1] = 100
	fmt.Println(myarray) //  [1 100 4 5]
	fmt.Println(array9)  // [1 3 4 5]
}

func main() {
	//定义定长数组
	var array1 = [2]int{} //表示数组长度为2,容量为2 内容为int 默认是[0 0]
	fmt.Println(array1)   //[0 0]
	fmt.Println(len(array1), cap(array1))
	array1 = [2]int{1, 2}
	fmt.Println(array1)

	var array2 = [3]string{"h", "e"}
	var array3 = [4]int{1, 2, 3}
	fmt.Println(array2) // [h e ] 注意这有个空字符
	fmt.Println(array3) //[1 2 3 0] int 默认是0

	var array4 = [5]int{0: 4, 4: 5} //指明一个元素为4，第五个元素为5
	fmt.Println(array4)             //[4 0 0 0 5]

	var array5 = [...]int{1, 3, 5, 7, 9}
	fmt.Println(array5)                             //[1 3 5 7 9]
	fmt.Printf("%v %v\n", len(array5), cap(array5)) //5 5

	var array6 = [5]int{0: 2, 1: 3, 2: 10, 3: 100, 4: 7}
	fmt.Println(array6)
	sum := 0
	for _, v := range array6 {
		sum = sum + v
	}
	fmt.Printf("sum=%v\n", sum)

	//二维数组
	// var  array7  [3][2]int

	// array7 = [3][2]int{
	// 	{1,2},
	// 	{3,4},
	// 	{5,6},
	// }

	// for  k, v ：=  range  array7 {
	// 	fmt.Print(v)
	// }
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
	//  多维数组只有第一层可以使用...来让编译器推导数组长度.数组是值类型，不可变.

	printArray(array9)
}

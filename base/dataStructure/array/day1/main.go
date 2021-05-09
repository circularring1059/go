package main

import "fmt"

func main() {
	//定义定长数组
	var array1 = [2]int{} //表示数组长度为2,容量为2 内容为int 默认是[0 0]
	fmt.Println(array1)   //[0 0]
	fmt.Println(len(array1), cap(array1))

	var array2 = [3]string{"h", "e"}
	var array3 = [4]int{1, 2, 3}
	fmt.Println(array2) // [h e ] 注意这有个空字符
	fmt.Println(array3) //[1 2 3 0] int 默认是0

	var array4 = [5]int{0: 4, 4: 5} //指明一个元素为4，第五个元素为5
	fmt.Println(array4)             //[4 0 0 0 5]

	var array5 = [...]int{1, 3, 5, 7, 9}
	fmt.Println(array5) //[1 3 5 7 9]
	fmt.Printf("%v %v\n", len(array5), cap(array5))

}

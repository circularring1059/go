package main

import "fmt"

func main() {
	fmt.Println(" to do your best")

	var slice1 []int
	// make  开辟空间
	slice1 = make([]int, 3, 5)

	slice1[2] = 2
	fmt.Println(slice1)

	// slice1[3] = 3 // error
	fmt.Println(slice1) //[0 0 2]
	//append  注意接收值还是原来的切片
	slice1 = append(slice1, 4)
	fmt.Println(slice1) //[0 0 2 4]

	//copy
	slice2 := []int{}
	copy(slice2, slice1)
	fmt.Println(slice2) // [] 没有开辟空间
	slice2 = make([]int, 5, 5)
	copy(slice2, slice1)
	fmt.Println(slice2) //[0 0 2 4 0]    注意最后后面的0

	var slice3 = []int{11, 22, 33, 44, 55, 66}
	copy(slice3, slice1)
	fmt.Println(slice3) //copy   把相同index 对应的值替换掉，目标切片的 length  要大于等于 源切片的length 才能有效copy

	//赋值
	var slice4 = slice1
	fmt.Println(slice4)

	slice1[0] = 100
	fmt.Println(slice4, slice1, slice2, slice3) //[100 0 2 4] [100 0 2 4] [0 0 2 4 0] [0 0 2 4 55 66]  赋值是应用，copy 是额外开辟内存空间的

	//切片元素删除，没有专门的方法，使用append
	var slice5 = []int{1, 2, 4, 5, 6, 7, 8}
	fmt.Println(slice5)
	//删除5 index=3
	slice5 = append(slice5[:3], slice5[4:]...)
	fmt.Println(slice5)

}

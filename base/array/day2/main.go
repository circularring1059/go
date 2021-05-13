package main

import "fmt"

func main() {
	fmt.Println("切片")
	// 声明切片
	var slice1 []int
	var slice2 []string
	var slice3 = []bool{true, false}

	fmt.Println(slice1, slice2, slice3) // [] [] [true false]
	fmt.Printf("长度：%v 容量： %v\n", len(slice3), cap(slice3))

	if slice1 == nil {
		fmt.Println("slice1  是空切片")
	}

	//初始化切片
	slice1 = []int{1, 2, 3, 4}
	slice2 = []string{"hello", "golang"}
	fmt.Println(slice1, slice2)

	//遍历切片
	for _, v := range slice2 {
		fmt.Println(v)
	}

	for i := 0; i < len(slice3); i++ {
		fmt.Println(slice3[i])
	}

	//切片长度与容量
	var slice4 = make([]int, 2, 5)                        //有两个元素(int 默认是0，现在可以放十个
	fmt.Println(slice4)                                   // [0 0]
	fmt.Printf("长度:%v 容量:%v\n", len(slice4), cap(slice4)) //长度:2 容量:5

	//声明一个  slice5 == nil
	var slice5 []int
	fmt.Printf("%v %d %d\n", slice5, len(slice5), cap(slice5)) // [] 0 0
	// slice5[0] = 3 // fmt.Printf("%v %d %d", slice5, len(slice5), cap(slice5))  没有容量不能赋值 runtime error: index out of range [0] with length 0
	slice5 = append(slice5, 7878)
	fmt.Printf("%v %d %d\n", slice5, len(slice5), cap(slice5)) //append  后自动开辟空间  [7878] 1 1

	slice5[0] = 9898    //修改值
	fmt.Println(slice5) //[9898]
}

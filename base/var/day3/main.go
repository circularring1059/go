package main

import "fmt"

func myfunc() (int, string) {
	return 100, "ring"
}

func changeA(b int) { //拷贝副本，值传递
	b++
	//fmt.Println(b) //11

}

func changeB(b *int) {
	*b++
}

func changeArray(c [3]int) {
	c[1] = 10
}

func changeArray1(c *[3]int) {
	(*c)[1] = 10
}

func changeSlice(d []string) {
	d = append(d, "长")
}

func changeSlice1(d *[]string) {
	(*d) = append((*d), "长")
}

func main() {
	number, _ := myfunc() // _  接收匿名变量，占位，不占内存
	fmt.Println(number)

	a := 10
	fmt.Println(a)
	changeA(a)
	fmt.Println(a)

	var b = 10
	fmt.Printf("b=%v\n", b)
	changeB(&b)
	fmt.Printf("b=%v\n", b)

	var array1 = [...]int{1, 2, 3}
	fmt.Println(array1)
	changeArray(array1)
	fmt.Println(array1)
	changeArray1(&array1)
	fmt.Println(array1)

	var slice1 = []string{}
	slice1 = make([]string, 2, 3)
	slice1[0] = "制"
	slice1[1] = "表"
	fmt.Println(slice1)
	changeSlice(slice1)
	fmt.Println(slice1)

	var slice2 = []string{"制", "表"}
	fmt.Println(slice2)
	changeSlice1(&slice2)
	fmt.Println(slice2)
	str1 := "dir"
	str2 := "file"
	fmt.Println(str1  + str2)

}

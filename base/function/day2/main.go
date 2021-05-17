package main

import "fmt"

func func1() {
	fmt.Println("this is a easy function")
}

func func2(x int) { // 一个int参数
	fmt.Println("int 参数")
	fmt.Println(x)
}

func func3(x int, y int) {
	fmt.Println(x + y)
}

func func4(x, y, z int, s string) {
	fmt.Println("多个同类型参数可以省略写")
	fmt.Println(x * y * z)
	fmt.Println(s)
}

func func5(s string, multilNu ...int) { //后面是类型的参数就可以传多个该类型的参数，这种省略写法只能放最后
	fmt.Println(s)
	fmt.Println(multilNu)
	fmt.Printf("%T\n", multilNu) //[]int slice 类型
}

func func6() string { //返回需要注明类型
	return "简单返回"
}

func func7() int {
	return 99
}

func func8() (res string) {
	res = "显示指定"
	return //显示指定 res return 后面可以省略不写
}

func func9() int { //返回一个int 类型的值
	var a = 1 + 3
	b := 3 + 1
	c := a + b
	return c

}

func func10() (str0, str1 string) {
	str0 = "fanhuiduoge"
	str1 = "返回多个"
	//return str0, str1
	return
}

func main() {
	fmt.Println("Learning function")
	func1()
	func2(9)
	func3(1, 1) //2
	func4(1, 2, 3, "str1")
	func5("str2", 0, 1, 3, 4, 5)
	fmt.Println(func6())
	fmt.Println(func7())
	fmt.Println(func8())
	fmt.Println(func9())
	fmt.Println(func10())
}

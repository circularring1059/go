package main

import "fmt"

/*其他玩法*/
//定义枚举
const (
	one   int = 1
	two   int = 2
	three int = 3
)

//配合iota  使用
const (
	a = iota // iota  初始值为0 a = 0
	b        // b = 1
	c        // c = 2
)

const (
	d, e = iota + 2, iota * 2 //d = 2 e = 0
	f, g                      //f =3 g =2

	h, k = iota * 3, iota + 1 // h = 6 k = 3

)

const (
	day1 = "day1"
	day2 //  day2="day1"
	day3 //  day3="day1"
)

func main() {
	//关键字 const  定义常量（只读属性）
	const name string = "google.com" //用双引号
	fmt.Println(name)

	//name = "baidul.com"   常量不可以被修改

	fmt.Println(one, two, three)

	fmt.Println(a, b, c)

	fmt.Println(d, e, f, g, h, k)

	fmt.Println(day1, day2, day3)
}

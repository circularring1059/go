//结构体在函数中值传递
package main

import "fmt"

type pet struct {
	name, gender string
	age          int
}

func changePetName(x pet) {
	x.name = "tom"
	fmt.Println(x) //{tom male 3}
}

//指针传值进行修改

func changePetAge(x *pet) {
	x.age = 4 //语法糖写法,和下面写法效果一样
	//(*x).age = 4 //
}

func main() {
	var com pet
	com.name, com.gender = "com", "male"
	com.age = 3
	// fmt.Println(com) //{com male 3}
	changePetName(com)
	fmt.Println(com) //{com male 3}未改变

	changePetAge(&com)
	fmt.Println(com) //age = 4
}

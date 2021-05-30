package main

import "fmt"

//结构嵌套

type nameAge struct {
	name string
	age  int
}
type person struct {
	// name string
	// age int
	na           nameAge
	idCardNumber int64
}

type cat struct {
	// name string
	// age int
	na    nameAge
	hobby []string
}

func main() {
	fmt.Println("结构体嵌套")
	var xiaoxiao person = person{
		na: nameAge{
			name: "xiaoxioa",
			age:  12,
		},
		idCardNumber: 123455,
	}

	tom := cat{
		na: nameAge{
			// name: "tom",
			// age:  2,
			"tom",
			2,
		},
		hobby: []string{"selpp", "eat"},
	}
	fmt.Println(xiaoxiao)
	fmt.Println(tom)
	fmt.Println(xiaoxiao.na.name)
	fmt.Println(tom.na.age)
}

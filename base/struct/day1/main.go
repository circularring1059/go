//struct  --->> 类  面向对象
package main

import "fmt"

type person struct {
	name, gender string
	age          int
	hobby        []string
}

func main() {
	var jack person
	fmt.Printf("%v %T\n", jack, jack) //{  0 []} main.person

	jack.name = "jack"
	jack.gender = "男"
	jack.age = 19
	jack.hobby = []string{"football", "basketball"}

	fmt.Println(jack) //{jack 男 19 [football basketball]}
	func(p person) {
		p.name = "tom" //值类型
		fmt.Println(p)
	}(jack)
	fmt.Println(jack)

	bobo := person{}
	fmt.Println(bobo)
	bobo.age = 50
	fmt.Println(bobo)
}

package main

import "fmt"

//  struct 初始化

type cat struct {
	name string
	age  int
}

func main() {
	//一
	var tom cat
	tom.name = "tom"
	tom.age = 1
	fmt.Printf("%v %v %v\n", tom.name, tom.age, tom)

	//二  cat 类型的指针
	var com = new(cat)
	(*com).name = "com"
	com.age = 1
	fmt.Printf("%T %v\n", com, com) //*main.cat &{com 1}
	fmt.Printf("%p\n", com)         //0xc000004090
	fmt.Printf("%p\n", &com)        //0xc0000d8020

	//三
	aom := cat{
		"aom",
		2,
	}
	fmt.Println(aom)
	fmt.Printf("%#v\n", aom) //main.cat{name:"aom", age:2}

	//四
	var bom = cat{
		name: "bom",
		age:  3,
	}
	fmt.Println(bom) //{bom 3}

	//五  cat  类型的指针
	var dom = &cat{
		name: "dom",
		age:  4,
	}
	fmt.Println(dom)
	fmt.Println(*dom)       //{dom 4}
	fmt.Printf("%T\n", dom) //*main.cat

}

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string //大写开头公有
	Age  int
}

func main() {
	fmt.Println("json 序列化")
	var man person = person{
		Name: "man",
		Age:  12,
	}
	fmt.Println(man)

	manJson, err := json.Marshal(man)
	if err != nil {
		fmt.Println("fail")
	} else {
		fmt.Println(string(manJson))
	}

}

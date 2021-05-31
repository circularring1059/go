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

	//序列化
	manJson, err := json.Marshal(man)
	if err != nil {
		fmt.Println("fail")
	} else {
		fmt.Println(string(manJson))
	}

	//反序列化
	str1 := `{"name": "circularring", "age": 13}`
	fmt.Println(str1)
	// var woman person
	var woman = new(person)
	json.Unmarshal([]byte(str1), &woman)
	fmt.Printf("%v\n", *woman)

	var middleMan person
	json.Unmarshal([]byte(str1), &middleMan)
	fmt.Println(middleMan)
}

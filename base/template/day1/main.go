package main

import (
	"fmt"
	"os"
	"text/template"
)

type person struct {
	Name string //开头要大写，否者访问不到
	Age  int
}

func main() {
	ring := person{"ring", 18}

	name := "ring"
	//模板
	temp := "my name is {{.}}\n"
	tmpl, err := template.New("test").Parse(temp)
	if err != nil {
		fmt.Println("template failed")
		return
	}
	err = tmpl.Execute(os.Stdout, name)
	if err != nil {
		fmt.Println("template execute  failed")
		return
	}
	temp1 := "my name is {{.Name}},  age is {{.Age}}"
	tmpl1, err := template.New("test").Parse(temp1)
	err = tmpl1.Execute(os.Stdout, ring)
	if err != nil {
		fmt.Println("template execute failed")
		fmt.Println(err)
		return
	}

}

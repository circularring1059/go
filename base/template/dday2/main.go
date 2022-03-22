package main

import (
	"fmt"
	"html/template"
	"os"
)

type person struct {
	Name string
	Age  int
}

func main() {
	fmt.Print("hello")

	ring := person{"ring", 18}

	tmpl, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Print("template failed")
		return
	}

	err = tmpl.Execute(os.Stdout, ring)
	if err != nil {
		fmt.Println("execute failed")
		return
	}
}

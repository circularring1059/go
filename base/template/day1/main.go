package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	name := "ring"
	//模板
	temp := "my name is {{.}}"
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

}

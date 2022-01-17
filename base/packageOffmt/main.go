package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	str1 = "ring"
	str2 = "yuan"
)

func bufioDemo() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}

func main() {
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./hello.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "好好学习，天天向上"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)

	str3 := fmt.Sprintln(str1, str2)
	str4 := fmt.Sprint(str1, str2)
	fmt.Println(str3) //ring yuan
	fmt.Println(str4) //ringyuan
	// fmtError := errors.New("出错了")
	// err := fmt.Errorf("这是一个错误%w", fmtError)
	// fmt.Println(err)
	e := errors.New("原始错误")
	w := fmt.Errorf("Wrap了一个错误%w", e)
	fmt.Println(w)

	//bufioDemo()

	osArg := os.Args        //[]string //go run main.go hello world
	fmt.Println("*", osArg) //[C:\Users\rice\AppData\Local\Temp\go-build608650344\b001\exe\main.exe hello wrorld \n]

	for i, v := range osArg {
		fmt.Println(i, v)
	}

}

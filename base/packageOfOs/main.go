package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed")
		return
	}
	defer file.Close()

	var tmp = make([]byte, 1024)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("file read failed")
		return
	}
	fmt.Println(n)
	fmt.Println(string(tmp[:n]))

}

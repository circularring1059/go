package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed")
		return
	}
	defer file.Close()

	//read  对象
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') // line type:string
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println("line")
			}
			fmt.Println("文件读取完毕")
			break
		}

		if err != nil {
			fmt.Println("file read fialed")
			return
		}
		fmt.Println("**")
		fmt.Print(line)
	}

}

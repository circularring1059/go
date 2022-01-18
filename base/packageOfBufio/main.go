package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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
		line, err := reader.ReadString('\n') // line type:string   一行一行读
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
		fmt.Print(line)
	}

	text, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("file read failed")
		return
	}
	fmt.Println(string(text))

}

func fileWrte() {
	file, err := os.OpenFile("./write.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "要写入的的text"
	file.Write([]byte(str))     //写入字节切片数据
	file.WriteString("content") //直接写入字符串数据
}

func bufioWrite() {
	file, err := os.OpenFile("./bufio.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("a line\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件

}

func ioutilWrite() {
	str := "content write on disk"
	err := ioutil.WriteFile("./iotuilWrte.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

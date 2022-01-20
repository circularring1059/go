package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	ret, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Print("request fialed")
		return
	}
	// 关闭连接
	defer ret.Body.Close()
	//获取response  content
	boy, err := ioutil.ReadAll(ret.Body) //好像比python requests 繁琐？？？
	if err != nil {
		fmt.Println("get response content fialed")
		return
	}
	fmt.Println(string(boy))
}

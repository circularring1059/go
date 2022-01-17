package main

import (
	"log"
)

func main() {
	//常规print 加上时间？
	// fmt.Println("study package of log")
	// log.Println("general log")
	// log.Printf("hello %s\n", "word")
	// log.Fatalln("fatal log")
	// log.Panicln("panic log")

	//设置日志格式
	// 控制输出日志信息的细节，不能控制输出的顺序和格式。
	// 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	// Ldate         = 1 << iota     // 日期：2009/01/23
	// Ltime                         // 时间：01:23:23
	// Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
	// Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
	// Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
	// LUTC                          // 使用UTC时间
	// LstdFlags     = Ldate | Ltime // 标准logger的初始值
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("print log") //2022/01/17 17:03:36.745682 C:/Users/rice/go/src/github.com/go/base/packageOfLog/main.go:26: print log
	log.Printf("%s is not supprot", "log")
	log.SetPrefix("[ring] ")
	log.Printf("%s is not supprot", "log")
}

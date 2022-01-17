package main

import (
	"fmt"
	"time"
)

//定时器
func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器   通道
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}

func main() {
	now := time.Now()
	fmt.Println(now.String()) //返回字符串类型
	fmt.Printf("%T\n", now)   //time.time
	fmt.Println(now)          //2022-01-07 14:24:37.1242977 +0800 CST m=+0.006823301
	fmt.Println(now.Unix())   //1641536705  时间戳
	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(now.Format("2006-01-02 15:04:05")) //2022-01-07 14:28:44
	fmt.Println(now.Format("2006/01/02 15:04:05")) //2022/01/07 14:28:44
	fmt.Println(now.Date())                        // Y M D
	fmt.Println(now.Clock())                       // H M S

	oneSecond := time.Second
	// Nanosecond  Duration = 1
	// Microsecond          = 1000 * Nanosecond
	// Millisecond          = 1000 * Microsecond
	// Second               = 1000 * Millisecond
	// Minute               = 60 * Second
	// Hour                 = 60 * Minute
	fmt.Printf("%T %v\n", oneSecond, oneSecond) //time.Duration 1s

	// time.Sleep(time.Second * 10) //sleep 10s
	fmt.Println("over")

	//时间加减
	fmt.Println(time.Now().Add(time.Minute * 3))
	fmt.Println(time.Now().AddDate(-1, 2, 10)) //  返回3 分钟后的时间  年 月 日

	starttime := time.Now()
	time.Sleep(time.Second * 2)
	endtime := time.Since(starttime)
	fmt.Println(starttime, endtime)
	fmt.Println("*")

	//设置时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	const shortForm = "2006-Jan-02"
	t1, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	t2, _ := time.ParseInLocation(shortForm, "1994-Nov-18", loc)
	fmt.Println(t1)
	fmt.Printf("%T -%v", t2, t2) //type time.time

	tickDemo()
}

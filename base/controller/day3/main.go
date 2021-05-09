package main

import "fmt"

func main() {
	//写法一
	var number1 int = 10
	if number1 == 10 {
		fmt.Println("number1=10")
	} else if number1 == 20 {
		fmt.Println("number1=120")
	} else if number1 == 30 {
		fmt.Println("number1=30")
	} else {
		fmt.Println("number1=?")
	}
	fmt.Println("number1:", number1)

	//写法二
	if number2 := 20; number2 == 10 {
		fmt.Println("number2=10")
	} else if number2 == 20 {
		fmt.Println("number2=20")
	} else if number2 == 30 {
		fmt.Println("number2=30")
	} else {
		fmt.Println("number2=?")
	}
	//fmt.Println("number2:", number2)  number2 只能判断中使用

	//for 循环
	var number3 int = 0
	for ; number3 < 10; number3++ {
		fmt.Print(number3) //0123456789
	}
	fmt.Println()
	fmt.Println(number3)

	for number4 := 0; number4 < 10; number4++ {
		fmt.Print(number4)
		fmt.Println()
	}
	// 同样number4 不能子啊该for循环外面打印

	number5 := 0
	for number5 < 10 {
		//fmt.Println("number5")
		fmt.Print(number5)
		number5++
	}

	// swith
	switch number4 := 7; number4 {
	case 1:
		fmt.Println("number4=1")
	case 2:
		fmt.Println("number4=2")
	case 7:
		fmt.Println("number4=7")
	case 5, 6, 3, 4:
		fmt.Println("numbe4 in 5,6,3,4")
	//case number4 > 2:  //error
	//	fmt.Println("number4 > 2")
	//case 2: 不能再写上面已有的
	// 	fmt.Print("")
	default:
		fmt.Println("default")
	}

	switch number4 := 7; {
	case number4 > 2:
		fmt.Println("number4 > 2")
		fallthrough // 加上fallthrouch 会执行下一个case,不管条件是否满足  兼顾c语言，用的不多
	case number4 > 18:
		fmt.Println("number4 > 18")
	default:
		fmt.Println("default")
	}

	//continue
	for number6 := 0; number6 < 10; number6++ {
		if number6 == 6 {
			fmt.Println("continue")
			continue
		}
		fmt.Println(number6)
	}

	//break
	for number6 := 0; number6 < 10; number6++ {
		if number6 == 6 {
			fmt.Println("break")
			break
		}
		fmt.Println(number6)
	}

	//双层循环时可以定义flag 跳出双循环
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			break
		}
	}
	// golang 里有个 goto 可以简化上述代码
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")

	//for range
	var str1 = "hell0"
	for i, v := range str1 {
		fmt.Println(i, v)     // v 是一个字符，打印出来是unicode 编码后的数字
		fmt.Printf("%c\n", v) // %c  显示该字符
	}

	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c\n", str1[i])
	}
}

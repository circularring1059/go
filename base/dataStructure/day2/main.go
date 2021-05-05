/*float32和float64。这两种浮点型数据格式遵循IEEE 754标准： float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32。
float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var number0 = math.MaxFloat32
	fmt.Println(number0)

	number1 := 1.234

	fmt.Println(number1)
	fmt.Printf("%T\n", number1) //默认为float64

	number2 := float32(1.24) //声明float32  的浮点数

	fmt.Printf("%T\n", number2)
}

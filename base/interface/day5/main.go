package main

import "fmt"

func getName(s interface{}) {
	fmt.Println(s)
}
func main() {
	fmt.Println("空接口")
	fmt.Println("多类型，函数参数")
	var map1 = map[string]interface{}{
		"one":   1,
		"tow":   "2",
		"three": [...]int{1, 2, 3},
		"four":  map[int]int{1: 1, 2: 2, 3: 3},
	}
	fmt.Println(map1)

	var array1 = [...]interface{}{1, "one", map[int]int{1: 1, 2: 2, 3: 3}, [...]int{1, 2, 3}}
	fmt.Println(array1)

	getName("getName")
	getName(123)

}

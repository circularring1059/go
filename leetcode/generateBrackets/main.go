package main

import "fmt"

func tools(m int, x int, y int, res string, ret *[]string) {
	if len(res) == 2*m {
		*ret = append(*ret, res)
	} else {
		if x < m {
			tools(m, x+1, y, res+"(", ret)
		}
		if y < x {
			tools(m, x, y+1, res+")", ret)
		}
	}

}

func generateBrackets(m int) []string {
	ret := make([]string, 0)
	res := ""
	tools(m, 0, 0, res, &ret)
	return ret
}

func main() {
	fmt.Println("start job")
	fmt.Println(generateBrackets(3))
}

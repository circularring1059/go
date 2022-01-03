package main

import "fmt"

func lengthOflongestSubstring(s string) int {
	strArry := []rune(s)
	var res int
	if len(strArry) == 1 {
		return 1
	}
	for i := 0; i < len(strArry); i++ {
		list := make(map[rune]int)
		list[strArry[i]] = i
		var max int
		for k := i + 1; k < len(strArry); k++ {
			if _, ok := list[strArry[k]]; ok {
				max = len(strArry[i:k])
				break
			} else {
				list[strArry[k]] = i
			}
			max = len(strArry[i : k+1])
		}
		if max > res {
			res = max
		}
	}
	return res
}

func main() {
	fmt.Println(lengthOflongestSubstring("ringyuanabcdefgjk"))
}

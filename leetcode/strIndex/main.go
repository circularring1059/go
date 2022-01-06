package main

import "fmt"

func strIndex(str1, str2 string) int {
	lengthOfstr2 := len(str2)
	lengthOfstr1 := len(str1)
	for i := 0; i <= lengthOfstr1-lengthOfstr2; i++ {
		flag := new(bool)
		for j := 0; j < lengthOfstr2; j++ {
			if str1[j+i] != str2[j] {
				*flag = true
				break
			}
		}
		if !*flag {
			return i
		}
	}
	return -1

}

func main() {
	// str1 := "ring"
	// fmt.Println(str1[1])
	// fmt.Printf("%c", str1[1])
	fmt.Println(strIndex("ring", "g"))
}

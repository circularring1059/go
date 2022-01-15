package main

import "fmt"

// func s(dig int) {
// 	// digArry = append(digArry, dig%10)
// 	// if dig/10 != 0 {
// 	// 	s(dig / 10)
// 	// }
// }

func oneTwo(dig int) []int {
	// fmt.Println(dig % 10)
	// if dig/10 != 0 {
	// 	oneTwo(dig / 10)
	// }
	digArry := make([]int, 0)
	s := func(a int) (c, d int) {
		// a := dig / 10
		// b := dig % 10
		return a / 10, a % 10

	}
	e, b := dig, 0
	for {

		e, b = s(e)
		digArry = append(digArry, b)
		if e == 0 {
			break
		}
	}
	return digArry

}
func main() {
	fmt.Println("one two three")
	fmt.Print(oneTwo(123))
}

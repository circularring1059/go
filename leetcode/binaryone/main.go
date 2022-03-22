package nain

import "fmt"

func binaryOne(n int) int {
	ret := 0
	m := n
	for {
		if m > 0 {
			m = m & (m - 1)
			ret++
		} else {
			break
		}
	}

	return ret
}

func main() {
	n := 1
	fmt.Println(binaryOne(n))
}

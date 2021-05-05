/* 复数 */
//复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位

package main

import "fmt"

func main() {

	var complex1 = 1 + 2i

	fmt.Println(complex1)
	fmt.Printf("%T\n", complex1) //complex128

	var complex2 = complex64(1 + 2i)
	fmt.Println(complex2)
}

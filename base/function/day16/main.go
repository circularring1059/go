package main

import "fmt"

func f1() int {
	x := 5
	x++
	defer func() {
		x++
		fmt.Println("defer", x)
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5

}

// func a() int {
// 	i := 0
// 	defer fmt.Println(i)
// 	i++
// 	return i
// }
// func c() (i int) {
// 	defer func() { i++ }()
// 	return 1
// }
// func f() int {
// 	i := 5
// 	defer func() {
// 		i++
// 	}()
// 	return i //5
// }

// func f1() (result int) {
// 	defer func() {
// 		result++
// 	}()
// 	return //1
// }

// func f2() (r int) {
// 	t := 5
// 	defer func() {
// 		t = t + 5
// 		fmt.Println(t)
// 	}()
// 	return t //5
// }

// func f3() (r int) {
// 	defer func(r int) {
// 		//r = r + 5 不规范
// 		r += 5
// 	}(r)
// 	return 1 //1
// }

func main() {
	fmt.Println("f4", f4())
	fmt.Println("f1", f1())
	fmt.Println("f2", f2())
	fmt.Println("f3", f3())

}

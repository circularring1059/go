package main

import "fmt"

func test1() {
	x, y := 10, 20

	defer func(i int) {
		println("defer:", i, y) // y 闭包引用  10 120
	}(x) // x 被复制

	x += 10
	y += 100
	println("x =", x, "y =", y) // x=20 y=120
}

func test(x int) {
	defer println("a")
	defer println("b")

	defer func() {
		println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
	}()

	defer println("c")
}

type person struct{ name string }

func (p *person) Close() {
	fmt.Println(p.name)
}

func main() {
	var a [5]struct{} //五个空的结构体
	fmt.Println(a)

	for i, v := range a {
		fmt.Println(i, v)
	}

	for i := range a {
		defer func() {
			fmt.Println(i) // 4 4 4 4 4
		}()
	}

	persons := []person{{"one"}, {"two"}, {"three"}}
	for _, v := range persons {
		//go语言并没有把这个明确写出来的this指针当作参数来看待
		defer v.Close() //three three three
		v1 := v
		defer v1.Close() //three two one
	}

	//test(0) //c b a  正常输出，然后panic  程序exit
	fmt.Println("*")
	test1()

	type n struct {
		name string
	}

	n1 := n{"ring"}
	fmt.Println(n1.name)
	func(s n) {
		s.name = "hello"
		fmt.Println(s.name)
	}(n1)
	fmt.Println(n1.name)

}

package main

import "fmt"

var array1 = [...]int{1, 2, 3, 4, 5, 6}

func main() {
	fmt.Println("try you best to do")

	//数组进行切片
	s1 := array1[0:2]
	fmt.Printf("%v-%d-%d\n", s1, len(s1), cap(s1)) // 长度为2    容量为6 左指针到数组最后

	s2 := array1[1:3]
	fmt.Printf("%v-%d-%d\n", s2, len(s2), cap(s2)) //长度为2  容量为5 左指针到数组最后

	//切片再切片
	s3 := s2[0:]
	fmt.Printf("%v-%d-%d\n", s3, len(s3), cap(s3)) //[2 3]-2-5

	s4 := s2[1:]
	fmt.Printf("%v-%d-%d\n", s4, len(s4), cap(s4)) //[3]-1-4

	array1[2] = 10
	fmt.Println(array1)
	fmt.Printf("%v-%d-%d\n", s3, len(s3), cap(s3)) //[2 10]-2-5
	fmt.Printf("%v-%d-%d\n", s4, len(s4), cap(s4)) //[10]-1-4

}

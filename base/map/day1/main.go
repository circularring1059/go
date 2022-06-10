package main

import "fmt"

//map 是没有顺序的
func main() {
	fmt.Println("Learning golang map")
	var map1 map[int]int
	fmt.Println(map1 == nil) //true
	var map2 map[string]int
	fmt.Println(map2 == nil) //true

	map1 = map[int]int{1: 1, 2: 2}
	fmt.Printf("%v-%v\n", len(map1), map1) //2-map[1:1 2:2]

	//map2["zero"] = 0 //panic: assignment to entry in nil map
	map2 = make(map[string]int, 1) //容量为5 容量不够时会动态扩容
	fmt.Println(map2)
	map2["one"] = 1
	fmt.Println(map2) //map[one:1]
	map2["two"] = 2
	fmt.Println(map2)
	map2["two"] = 22  //修改value 值
	fmt.Println(map2) //map[one:1 two:22]

	//遍历map
	for k, v := range map2 {
		fmt.Printf("%v-%v\n", k, v)
	}

	for v := range map2 { //只接收key
		fmt.Printf("%v\n", v)
	}

	for _, v := range map2 {
		fmt.Printf("%v\n", v)
	}

	//不存在的key
	fmt.Println(map2["three"]) //int 对应返回 0
	//做判断
	v, ok := map2["three"]
	if !ok {
		fmt.Println("没有此key")
	} else {
		fmt.Print(map2["three"], v)
	}

	//delete  key
	delete(map2, "one")

	fmt.Println(map2)

	map3 := make(map[int]string) //可以不写容量
	map3[1] = "c++"
	map3[2] = "python"
	map3[3] = "golang"
	map3[4] = "java"
	fmt.Println(map3)

	map4 := map[int]int{1: 1, 2: 2}
	fmt.Println(map4)
	for k, v := range map4 {
		fmt.Println(k, v)
	}
	changeMap(map4)
	fmt.Println(map4)  //引用传递  值改变

	map5 := map[int]int{1:1, 2:2, 3:3, 4:4}
	fmt.Println(len(map5))
	slice1 := make([]int, 0)
	for k, _:= range map5{
		fmt.Println(k, v)
		slice1 = append(slice1, k)
		fmt.Println(slice1)
	}

	for _, v := range slice1{
		fmt.Println(v, map5[v])
	}

	// map6 := make(map[int]int, 0)
	// map6[1] = 1
	// fmt.Println(map6)
	
}

func changeMap(m map[int]int) {
	m[1] = 11
	fmt.Println(m)
}

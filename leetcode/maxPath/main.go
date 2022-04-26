package main 

import "fmt"


func main (){
	fmt.Println("start")
	fmt.Println(maxPath(1, 1))
}



func backtack(x, y int, maxPath *int){
	if x == 0  && y ==0 {
		*maxPath ++
	}else {
		if x >0 {
			backtack(x-1, y, maxPath)
		}
		if y > 0 {
			backtack(x, y-1, maxPath)
		}
	}

}

func maxPath(x, y int) int{
	maxPath := 0
	backtack(x, y, &maxPath)
	return maxPath

}
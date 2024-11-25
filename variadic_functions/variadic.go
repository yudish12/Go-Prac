package main

import "fmt"

//variadic functions are functions which can take variable number of parameters (fixed type or any type)
// we cannot return varidaic

func variadic(nums ...int)int{ // for any type use func variadic(vals ...interface{})
	total:=0

	for _,val:=range nums{
		total+=val
	}
	return total
}

func main(){
	sum := variadic(1,2,3,4,6)
	fmt.Println(sum)
}
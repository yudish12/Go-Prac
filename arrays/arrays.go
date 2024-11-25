package main

import (
	"fmt"
	sliceLocal "hello-world/arrays/slices"
)

func main(){
	fmt.Println("===Array related stuff in this package===")

	arr := [4]string{"hi","how","are","you"} //fully-intitialized
	
	n := len(arr)

	for i:=0;i<n;i++ {
		fmt.Println(arr[i])
	}

	//array initialized with default values as "",0,false,0.0
	arr2 := [2]string {} //not initialized
	fmt.Println(arr2)

	//intialize specific elements
	arr3 := [5]int{2:15,4:11}
	fmt.Println(arr3)

	//2d arrays
	// matrix := [3][3]int{{1,2,3},{1,2,3},{1,2,3}}
	matrix := [3][2]int{{1,2}} //partially initialized 2d array
	fmt.Println(matrix)
	rows := len(matrix)
	cols := len(matrix[0])
	fmt.Println(rows,cols)

	sliceLocal.SlicesEx()

}
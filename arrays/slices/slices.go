package slices

import (
	"fmt"
	"slices"
)

func SlicesEx(){
	//initialize slice
	fmt.Println("===Slice Related stuff===")
	nums := []int{}
	fmt.Println(nums==nil) //empty slice is not nil

	//initialize slice with a initital length using make function
	sliceArr := make([]int,2) //we can add more number of elements than 2 but initially slice will be of 2 length
	fmt.Println(len(sliceArr),sliceArr)

	//capacity
	fmt.Println(cap(sliceArr)) //initial capacity

	sliceArr = append(sliceArr, 3)
	sliceArr = append(sliceArr, 4)
	sliceArr = append(sliceArr, 5)

	fmt.Println(len(sliceArr),cap(sliceArr)) //capacity increases by twice and length increases normally

	//copy function in slices
	nums1 := make([]int,3)
	nums2 := []int{1,2,3}

	copy(nums1,nums2) //copies elements from nums2 to nums1 note if nums1 length is 0 it won't copy anything
	fmt.Println(nums1)

	//slice operator
	sliceOp := []int{1,2,3,4}
	
	fmt.Println(sliceOp[0:3])
	fmt.Println(sliceOp[1:3])
	fmt.Println(sliceOp[:2])
	fmt.Println(sliceOp[2:])
	fmt.Println(sliceOp[:])

	//slice pacakge
	slicePkg := make([]int,4)
	copy(slicePkg,sliceOp)
	// slicePkg = append(slicePkg, 5)

	fmt.Println(slices.Equal(slicePkg,sliceOp))
	fmt.Println(slices.Chunk(slicePkg,2))

	for i := range slices.Chunk(slicePkg,2){
		fmt.Println(i)
	}
	
	//2d slices
	matrixSlice := make([][]int, 3)
	for i,val := range matrixSlice {
		matrixSlice[i] = make([]int, 2)
		fmt.Println(val)
	}

	fmt.Println(matrixSlice)

}


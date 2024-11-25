package main

import (
	"fmt"
	"strconv"
)

func Btoi(b bool) int {
    if b {
        return 1
    }
    return 0
 }

func main(){
	//numeric conversion to string
	i := 10
	str := strconv.Itoa(i)
	fmt.Printf("data type of str is %T\n",str)

	//fmt to print number with string using printf
	fmt.Printf("This is the string value of ten: %d\n", i)

	//float conversion
	fval := float64(i)

	fmt.Printf("data type of fval is %T\n",fval)

	fval += 1.231231312312
	fmt.Println(fval)

	//float to string
	fstr := strconv.FormatFloat(fval,'g', -1, 64)
	fmt.Println(fstr)


	//number to bool and vice versa
	n := 10
	b := n!=0
	fmt.Println(b)

	n2:=Btoi(b)
	fmt.Println(n2)

}
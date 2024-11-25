package main

import "fmt"

func main(){
	//range gives key,value; index,value in case of maps and slices,arrays respectively it also works with strings and channels
	var m map[string]string = make(map[string]string)
	m["lname"] = "fname"
	m["age"] = "years"

	arr := make([]int,4)

	//range with map
	for key,val := range m{
		fmt.Println(key,val)
	}

	//range with slice/array
	for i,val := range arr{
		fmt.Println(i,val)
	}

	//range with strings gives rune i is starting byte and val is ascii code (more in runes directory)
	str:="golang"
	for i,val:=range str{
		fmt.Println(i,val)
	}
}
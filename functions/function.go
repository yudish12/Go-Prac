package main

import (
	"fmt"
	"strconv"
)

func add(a int, b int) int {
	return a + b
}

//In go functions can return multiple values also using below syntax
//here 3 values of type strings should be returned
func getLangs()(string,string,string){
	return "javascript","python","rust"
}

func functionParam (fn func(a int,b string)map[string]int, key string,a int, b string)string{
	m := fn(a,b)
	val,ok := m[key]

	if ok {
		return "found"+ " "+ strconv.Itoa(val)
	}else {
		return "not found"
	}
}

func makeMap(a int, b string)map[string]int{
	m:= make(map[string]int)
	m[b]=a
	return m
}

func main() {
	res := add(1,3)
	lang1,lang2,lang3 := getLangs() // this is how we can desctructure to get all 3 return values
	fmt.Println(res,lang1,lang2,lang3)

	fmt.Println(functionParam(makeMap,"num",2,"nums"))

}
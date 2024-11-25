package main

import (
	"fmt"
	"maps"
)

func main(){
	//create a map
	m := map[string]string{"hi":"hello"}	
	//create map using make function
	m2 := make(map[string]string,3)

	fmt.Println(m,m2)

	//accessing elements etc in maps
	fmt.Println(m["hi"],m2["hi"]) //if key is not present in map it returns zero value

	//string and int map
	m3 := make(map[string]int)
	m3["phone"] = 2
	m3["key"] = 3
	m3["age"] = 23

	fmt.Println(m3,m3["age"])
	//changing value in map for a key
	m3["key"] = 5
	fmt.Println(m3["key"])

	//length of map
	fmt.Println(len(m3))

	//deleting an element in map
	delete(m3,"key")
	fmt.Println(m3)

	val,ok := m3["key"] 


	fmt.Println(val,ok)
	//if ok is true it means the element is present in map with specified key and the value is val 
	// if ok is false element with the key is not there in map and value val is zero value
	if ok {
		fmt.Println(m3["key"])
	}else {
		fmt.Println(val) //gives the 0 value
	}

	fmt.Println(maps.Equal(m,m2))

}
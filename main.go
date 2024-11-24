package main

import (
	"fmt"
	"hello-world/utils"
)

func main(){
	fmt.Println("Hello world")
	var name string = "yudish"
	name2 := "arrr"
	fmt.Println("Hello "+ name+" "+ name2)
	utils.Utility()
	fmt.Println(utils.PublicVar)
	utils.Help()
}
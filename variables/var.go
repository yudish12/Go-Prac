package main

import (
	"fmt"
	"hello-world/variables/contants"
	"strings"
)

func main(){
	//int variables

	//3 ways to declare a variable
	var a int = 10
	b := 10
	var c = 10
	fmt.Println(a,b,c)


	//float variables
	var fa float32 = 23.22 //float64 also exists
	fb := 221.1212312
	fmt.Println(fa,fb)

	// booleans 
	boola := true
	fmt.Println(boola)

	//string variables
	var str string = "Hello World"
	arr := strings.Split(str, " ")
	fmt.Println(arr[1])

	contants.ContantExamples();

}
package main

import (
	"fmt"
	"hello-world/loops-tut/loops"
)

func ExecLoops(){
	fmt.Println("===Loops Tutorial===")
	loops.ForLoop()	
	loops.WhileLoop()
}

func main(){
	ExecLoops()
}
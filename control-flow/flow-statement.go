package main

import (
	"fmt"
	ifelse "hello-world/control-flow/ifElse"
	"hello-world/control-flow/switchcase"
)

func main(){
	ifelse.IfElse(5)
	fmt.Println(switchcase.SwitchCase(3))
	switchcase.Typeswitch(true)
}
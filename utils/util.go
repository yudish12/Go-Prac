package utils

import "fmt"

var PublicVar = "Public"
var privateVar = "Private"

func Utility() {
	fmt.Println("Utility Called", privateVar)
}
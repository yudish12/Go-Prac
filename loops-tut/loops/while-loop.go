package loops

import (
	"fmt"
	"strconv"
)

func WhileLoop(){
	fmt.Println("\nRunning while loop")
	var i int
	// for keyword instead of while diff from other languages
	for i < 10 {
		fmt.Println(strconv.Itoa(i) + " index")
		i++
	}
}
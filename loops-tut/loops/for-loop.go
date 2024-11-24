package loops

import (
	"fmt"
)

func ForLoop(){
	fmt.Println("Running for loop")
	var i int32
	for i = 0;i<10;i++{
		fmt.Println(i)
	}
}
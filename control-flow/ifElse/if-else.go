package ifelse

import "fmt"

func IfElse(i int32){

	if(i==10){
		fmt.Println("Is ten")
	}else if (i==5){
		fmt.Println("Is five")
	}else {
		fmt.Println("Unrecognized value")
	}

}
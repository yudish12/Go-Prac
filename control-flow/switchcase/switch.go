package switchcase

import (
	"fmt"
)

func SwitchCase(val int32)string {
	switch val {
		case 1:
			return "one"
		case 2:
			return "two"
		case 3:
			return "three"
		case 4:
			return "four"
		case 5:
			return "five"
	}
	return "other"
}

func Typeswitch(i interface {}){
	switch i.(type) {
	case int:
		fmt.Println("int value")
	case string:
		fmt.Println("string value")
	case float32:
		fmt.Println("float32 value")
	case bool:
		fmt.Println("Boolean value")
	}
}
package contants

import "fmt"

func ContantExamples() {
	const pi = 3.14
	// pi = 12312 this will give error as it is a contant
	fmt.Println(pi)
}
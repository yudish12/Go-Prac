package main

import "fmt"

func counter()func()int{
	count := 0

	return func() int {
		count++
		return count
	}
}

func main(){
	incr := counter()
	fmt.Println(incr())
	fmt.Println(incr())
	fmt.Println(incr())
}
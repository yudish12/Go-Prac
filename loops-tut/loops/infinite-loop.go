package loops

import (
	"fmt"
	"time"
)


func InfiniteLoop(){
	breakLoop := false
	i:=0
	for {
		fmt.Println(i)
		i++
		time.AfterFunc(3*time.Second, func() {
			breakLoop=true
		})
		if(breakLoop) {
			break
		}
	}
}
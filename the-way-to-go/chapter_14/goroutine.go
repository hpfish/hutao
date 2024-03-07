package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)

	for {
		select {
		case a := <-tick:
			fmt.Println(a)
		}
	}

}

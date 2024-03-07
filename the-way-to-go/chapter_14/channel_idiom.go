package main

import (
	"fmt"
	"time"
)

func main() {
	stream := pump2()
	go suck2(stream)
	time.Sleep(1e9)
}

func pump2() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck2(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

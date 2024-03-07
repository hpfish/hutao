package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)
	ch := make(chan int, 1)
	go func() {
		for {
			ch <- 1
		}
	}()

	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

	for {
		select {
		case <-ch:
			fmt.Println("get value")
		case <-time.After(2 * time.Second):
			fmt.Println("time after")
			return
		case <-timeout:
			fmt.Println("timeout")
			return
		}
	}
}

package main

import "fmt"

//func main() {
//	defer println("in main")
//	go func() {
//		defer println("in goroutine")
//		panic("")
//	}()
//	time.Sleep(1 * time.Second)
//}

//func main() {
//	defer fmt.Println("in main")
//	if err := recover(); err != nil {
//		fmt.Println(err)
//	}
//
//	panic("unknown err")
//}

func main() {
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")
}

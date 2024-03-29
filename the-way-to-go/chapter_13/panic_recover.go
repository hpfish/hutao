package main

import "fmt"

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		fmt.Println("before recover")
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n")

}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}

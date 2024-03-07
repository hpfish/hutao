package main

import "fmt"

func foo2() *int {
	t := 3
	return &t

}

func main() {
	x := foo2()
	fmt.Println(*x)
}

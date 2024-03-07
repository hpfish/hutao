package main

import "fmt"

var s string

func main() {
	a()
	b()
}

func a() {
	s = "a"
	fmt.Println(s)
}

func b() {
	s = "b"
	fmt.Println(s)
}

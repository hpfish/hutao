package main

import "fmt"

func main() {
	var p *[]int = new([]int) // *p == nil; with len and cap 0
	fmt.Println(*p)
	var v []int = make([]int, 10, 50)
	fmt.Println(&v)
}

package main

import "fmt"

func main() {
	var s = []int{5}
	fmt.Printf("%+v\n", s)
	s = append(s, 7)
	fmt.Printf("%+v\n", s)
	x := append(s, 9)
	fmt.Println(x)
	y := append(s, 10)
	fmt.Println(y)
}

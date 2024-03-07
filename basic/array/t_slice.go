package main

import "fmt"

func main() {
	arr1 := new([2]string)
	slice1 := make([]string, 2)
	fmt.Printf("arr1: %v \n", arr1)
	fmt.Printf("slice1: %v\n", slice1)

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := slice[2:6:7]
	fmt.Println(cap(s1))
	fmt.Println(cap(s2))
}

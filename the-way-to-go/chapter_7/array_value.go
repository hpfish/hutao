package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := arr1
	arr1[1] = 5
	fmt.Printf("arr1: %v \n", arr1)
	fmt.Printf("arr2: %v \n", arr2)

}

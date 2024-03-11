package main

import "fmt"

func main() {
	arr := []int{1}
	fmt.Printf("arr: %p\n", &arr)
	myfunc1(arr)
	fmt.Println(arr)

	arr = append(arr, 3)
	fmt.Printf("arr2: %p\n", &arr)
	arr = append(arr, 4)
	fmt.Printf("arr3: %p\n", &arr)
	myfunc2(arr)
	fmt.Println(arr)
}

func myfunc1(arr []int) {
	fmt.Printf("before: %p\n", &arr)
	arr = append(arr, 2)
	fmt.Printf("after: %p\n", &arr)
	fmt.Println()
	arr[0] = 0
	fmt.Println(arr)
	return
}

func myfunc2(arr []int) {
	arr = append(arr, 5)
	fmt.Printf("after: %p\n", &arr)
	arr[0] = 9
	fmt.Println(arr)
	return
}

package main

import (
	"fmt"
)

func main() {
	arr := []int{1}
	fmt.Println("cap:", cap(arr))
	myfunc1(arr)
	fmt.Println(arr)

	arr = append(arr, 3)
	fmt.Println("cap:", cap(arr))
	arr = append(arr, 4)
	fmt.Println("cap:", cap(arr))
	myfunc2(arr)
	fmt.Println(arr)
	fmt.Println("cap:", cap(arr))
}

func myfunc1(arr []int) {
	arr = append(arr, 2)
	fmt.Println("cap:", cap(arr))
	arr[0] = 0
	fmt.Println(arr)
	return
}

func myfunc2(arr []int) {
	arr = append(arr, 5)
	fmt.Println("cap:", cap(arr))
	arr[0] = 9
	fmt.Println(arr)
	fmt.Println("cap:", cap(arr))
	return
}

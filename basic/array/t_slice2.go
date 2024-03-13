package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// 去掉最后一位
	fmt.Println(s[:len(s)-1])
	fmt.Println(s[3:])
}

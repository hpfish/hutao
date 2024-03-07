package main

import "fmt"

func main() {
	var a = "abcdefg"
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	//fmt.Println(&a[0])
}

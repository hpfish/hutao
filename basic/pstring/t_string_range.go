package main

import "fmt"

func main() {
	var s = "abcde"
	for i, c := range s {
		fmt.Println(i, ":", string(c))
	}
}

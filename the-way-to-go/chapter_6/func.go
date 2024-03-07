package main

import (
	"fmt"
	"strings"
)

//func flushICache(begin, end uintptr)

type binOp func(int, int) int

func main() {
	addBmp := MakeAddSuffix(".bmp")
	fmt.Println(addBmp("test"))
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

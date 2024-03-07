package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "I am handsome."
	for _, val := range strings.Fields(str) {
		fmt.Printf("%s - ", val)
	}
}

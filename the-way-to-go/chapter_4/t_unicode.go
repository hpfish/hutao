package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(unicode.IsDigit('A'))
	fmt.Println(unicode.IsLetter('A'))
	fmt.Println(unicode.IsSpace(' '))
}

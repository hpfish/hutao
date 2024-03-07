package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	buf.WriteString("a")
	buf.WriteString("b")
	buf.WriteString("c")
	a1, a2 := buf.Bytes()[:1], buf.Bytes()[1:]
	fmt.Println(a1)
	fmt.Println(a2)
}

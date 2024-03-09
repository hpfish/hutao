package main

import "fmt"

func main() {
	var stack []int

	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	stack = append(stack, 5)
	stack = append(stack, 4)

	peek := stack[len(stack)-1]
	fmt.Println(peek)

	pop := stack[len(stack)-1]
	fmt.Println(pop)
	stack = stack[:len(stack)-1]

	size := len(stack)
	fmt.Println(size)

	isEmpty := len(stack) == 0
	fmt.Println(isEmpty)
}

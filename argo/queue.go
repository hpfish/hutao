package main

import (
	"container/list"
	"fmt"
)

func main() {
	queue := list.New()

	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)
	queue.PushBack(5)
	queue.PushBack(4)

	peek := queue.Front()
	fmt.Println(peek)

	pop := queue.Front()
	queue.Remove(pop)
	fmt.Println(pop)
	fmt.Println(queue)

	fmt.Println(queue.Len())
	fmt.Println(queue.Len() == 0)
}

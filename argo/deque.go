package main

import (
	"container/list"
	"fmt"
)

func main() {
	deque := list.New()

	deque.PushBack(2)
	deque.PushBack(5)
	deque.PushBack(4)
	deque.PushFront(3)
	deque.PushFront(1)

	front := deque.Front()
	fmt.Println(front.Value)
	rear := deque.Back()
	fmt.Println(rear.Value)

	v := deque.Remove(front)
	fmt.Println(v)
	v = deque.Remove(rear)

	size := deque.Len()
	fmt.Println(size)
}

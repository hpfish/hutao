package main

import (
	"container/list"
	"fmt"
)

type linkedListStack struct {
	data *list.List
}

func newLinkedListStack() *linkedListStack {
	return &linkedListStack{
		data: list.New(),
	}
}

func (s *linkedListStack) push(value int) {
	s.data.PushBack(value)
}

func (s *linkedListStack) pop() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	s.data.Remove(e)
	return e.Value
}

func (s *linkedListStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	return e
}

func (s *linkedListStack) size() int {
	return s.data.Len()
}

func (s *linkedListStack) isEmpty() bool {
	return s.data.Len() == 0
}

func (s *linkedListStack) toList() *list.List {
	return s.data
}

func main() {
	s := newLinkedListStack()
	s.push(1)
	s.push(2)
	s.push(4)
	s.push(3)

	fmt.Println(s.toList())

	pop := s.pop()
	fmt.Println(pop)
	fmt.Println(s.toList())

	peek := s.peek()
	fmt.Println(peek)
	peek = s.peek()
	fmt.Println(peek)

}

package main

import "container/list"

type linkedListDeque struct {
	data *list.List
}

func newLinkedListDeque() *linkedListDeque {
	return &linkedListDeque{data: list.New()}
}

func (s *linkedListDeque) pushFirst(value any) {
	s.data.PushFront(value)
}

func (s *linkedListDeque) pushLast(value any) {
	s.data.PushBack(value)
}

func (s *linkedListDeque) popFirst(value any) any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Front()
	s.data.Remove(e)
	return e.Value
}

func (s *linkedListDeque) popLast() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	s.data.Remove(e)
	return e.Value
}

func (s *linkedListDeque) peekFirst() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Front()
	return e.Value
}

func (s *linkedListDeque) peekLast() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	return e.Value
}

func (s *linkedListDeque) isEmpty() bool {
	return s.data.Len() == 0
}

func (s *linkedListDeque) toList() *list.List {
	return s.data
}

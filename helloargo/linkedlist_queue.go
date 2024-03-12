package main

import "container/list"

type linkedListQueue struct {
	data *list.List
}

func newLinkedListQueue() *linkedListQueue {
	return &linkedListQueue{
		data: list.New(),
	}
}

func (s linkedListQueue) push(value any) {
	s.data.PushBack(value)
}

func (s *linkedListQueue) pop() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Front()
	s.data.Remove(e)
	return e.Value
}

func (s *linkedListQueue) peek() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Front()
	return e.Value
}

/* 获取队列的长度 */
func (s *linkedListQueue) size() int {
	return s.data.Len()
}

/* 判断队列是否为空 */
func (s *linkedListQueue) isEmpty() bool {
	return s.data.Len() == 0
}

/* 获取 List 用于打印 */
func (s *linkedListQueue) toList() *list.List {
	return s.data
}

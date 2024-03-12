package main

import "fmt"

// 基于环形数组实现的队列
type arrayQueue struct {
	nums        []int // 用于存储队列元素的数组
	front       int   // 队首指针， 指向队首元素
	queSize     int   // 队列长度
	queCapacity int   // 队列容量（即最大容纳元素数量）
}

func newArrayQueue(queCapacity int) *arrayQueue {
	return &arrayQueue{
		nums:        make([]int, queCapacity),
		front:       0,
		queSize:     0,
		queCapacity: queCapacity,
	}
}

func (q *arrayQueue) size() int {
	return q.queSize
}

func (q *arrayQueue) isEmpty() bool {
	return q.queSize == 0
}

func (q *arrayQueue) push(num int) {
	if q.queSize == q.queCapacity {
		return
	}
	rear := (q.front + q.size()) % q.queCapacity
	fmt.Println("num", num, "front", q.front, "size", q.queSize, "queCapacity", q.queCapacity, "rear", rear)
	q.nums[rear] = num
	q.queSize++

}

func (q *arrayQueue) peek() any {
	if q.isEmpty() {
		return nil
	}
	return q.nums[q.front]
}

func (q *arrayQueue) pop() any {
	num := q.peek()
	q.front = (q.front + 1) % q.queCapacity
	q.queSize--
	return num
}

func (q *arrayQueue) toSlice() []int {
	rear := (q.front + q.queSize)
	if rear >= q.queCapacity {
		rear %= q.queCapacity
		return append(q.nums[q.front:], q.nums[:rear]...)
	}
	return q.nums[q.front:rear]
}

func main() {
	queue := newArrayQueue(2)
	queue.push(1)
	queue.push(2)
	p := queue.pop()
	fmt.Println(p)
	p = queue.pop()
	fmt.Println(p)
	queue.push(3)
}

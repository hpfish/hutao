package main

import "fmt"

type arrayDeque struct {
	nums        []int // 用于存储双向队列元素的数组
	front       int   // 队首指针，指向队首元素
	queSize     int   // 双向队列长度
	queCapacity int   // 队列容量（即最大容纳元素数量）
}

func newArrayDeque(queCapacity int) *arrayDeque {
	return &arrayDeque{
		nums:        make([]int, queCapacity),
		queCapacity: queCapacity,
		queSize:     0,
		front:       0,
	}
}

func (q *arrayDeque) size() int {
	return q.queSize
}

func (q *arrayDeque) isEmpty() bool {
	return q.queSize == 0
}

func (q *arrayDeque) index(i int) int {
	// 通过取余操作实现数据首尾相连
	// 当 i 越过数组尾部后， 回到头部
	// 当 i 越过数组头部后， 回到尾部
	return (i + q.queCapacity) % q.queCapacity
}

// 队首入队
func (q *arrayDeque) pushFirst(num int) {
	if q.queSize == q.queCapacity {
		fmt.Println("双向队列已满")
		return
	}
	// 队首指针向左移动一位
	// 通过取余操作实现 front 越过数组头部后回到尾部
	q.front = q.index(q.front - 1)
	q.nums[q.front] = num
	q.queSize++
}

// 队尾入队
func (q *arrayDeque) pushLast(num int) {
	if q.queSize == q.queCapacity {
		fmt.Println("双向队列已满")
		return
	}
	// 计算队尾指针，指向队尾索引 + 1
	rear := q.index(q.front + q.queSize)
	q.nums[rear] = num
	q.queSize++
}

func (q *arrayDeque) popFirst() any {
	num := q.peekFirst()
	q.front = q.index(q.front + 1)
	q.queSize--
	return num
}

func (q *arrayDeque) popLast() any {
	num := q.peekLast()
	q.queSize--
	return num
}

func (q *arrayDeque) peekFirst() any {
	if q.isEmpty() {
		return nil
	}
	return q.nums[q.front]
}

func (q *arrayDeque) peekLast() any {
	if q.isEmpty() {
		return nil
	}
	last := q.index(q.front + q.queSize - 1)
	return q.nums[last]
}

func (q *arrayDeque) toSlice() []int {
	res := make([]int, q.queSize)
	for i, j := 0, q.front; i < q.queSize; i++ {
		res[i] = q.nums[q.index(j)]
		j++
	}
	return res
}

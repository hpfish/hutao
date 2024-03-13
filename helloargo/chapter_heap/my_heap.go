package chapter_heap

import "fmt"

type maxHeap struct {
	data []any
}

func newMaxHeap(nums []any) *maxHeap {
	h := &maxHeap{data: nums}
	for i := h.parent(len(h.data) - 1); i >= 0; i-- {
		h.siftDown(i)
	}
	return h
}

func (h *maxHeap) size() int {
	return len(h.data)
}

func (h *maxHeap) isEmpty() bool {
	return len(h.data) == 0
}

func (h *maxHeap) left(i int) int {
	return 2*i + 1
}

func (h *maxHeap) right(i int) int {
	return 2*i + 2
}

func (h *maxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *maxHeap) peek() any {
	return h.data[0]
}

func (h *maxHeap) push(val any) {
	h.data = append(h.data, val)
	// 从底至顶堆化
	h.siftUp(len(h.data) - 1)
}

func (h *maxHeap) swap(i, p int) {
	tmp := h.data[i]
	h.data[i] = h.data[p]
	h.data[p] = tmp
}

// 从节点 i 开始， 从底至顶堆化
func (h *maxHeap) siftUp(i int) {
	for true {
		p := h.parent(i)
		if p < 0 || h.data[i].(int) <= h.data[p].(int) {
			break
		}
		h.swap(i, p)
		i = p
	}
}

func (h *maxHeap) pop() any {
	if h.isEmpty() {
		fmt.Println("error")
		return nil
	}

	h.swap(0, h.size()-1)
	val := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.siftDown(0)
	return val
}

// 从节点 i 开始， 从顶层至底堆化
func (h *maxHeap) siftDown(i int) {
	for true {
		l, r, max := h.left(i), h.right(i), i
		if l < h.size() && h.data[l].(int) > h.data[max].(int) {
			max = l
		}
		if r < h.size() && h.data[r].(int) > h.data[max].(int) {
			max = r
		}
		if max == i {
			break
		}
		h.swap(i, max)
		i = max
	}
}

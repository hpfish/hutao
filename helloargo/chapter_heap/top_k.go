package chapter_heap

import "container/heap"

type minHeap []any

func (h *minHeap) Len() int           { return len(*h) }
func (h *minHeap) Less(i, j int) bool { return (*h)[i].(int) < (*h)[j].(int) }
func (h *minHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

// Push heap.Interface 的方法，实现推入元素到堆
func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Pop heap.Interface 的方法，实现弹出堆顶元素
func (h *minHeap) Pop() any {
	// 待出堆元素存放在最后
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

// Top 获取堆顶元素
func (h *minHeap) Top() any {
	return (*h)[0]
}

func topKHeap(nums []int, k int) *minHeap {
	h := &minHeap{}
	heap.Init(h)
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}
	for i := k; i < len(nums); i++ {
		if nums[i] > h.Top().(int) {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}
	return h
}

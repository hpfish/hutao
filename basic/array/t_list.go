package main

import "fmt"

type myList struct {
	arrCap      int
	arr         []int
	arrSize     int
	extendRatio int
}

func newMyList() *myList {
	return &myList{
		arrCap:      10,              // 列表容量
		arr:         make([]int, 10), // 数组（存储列表元素）
		arrSize:     0,               // 列表长度（当前元素数量）
		extendRatio: 2,               // 每次列表扩容的倍数
	}
}

func (l *myList) size() int {
	return l.arrSize
}

func (l *myList) capacity() int {
	return l.arrCap
}

func (l *myList) get(index int) int {
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	return l.arr[index]
}

func (l *myList) set(num, index int) {
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	l.arr[index] = num
	return
}

func (l *myList) add(num int) {
	if l.arrSize > l.arrCap {
		l.extendCapacity()
	}
	l.arr[l.arrSize] = num
	l.arrSize++
	return
}

func (l *myList) insert(num, index int) {
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	if l.arrSize == l.arrCap {
		l.extendCapacity()
	}
	// 将索引 index 以及之后的元素都向后移动一位
	for j := l.arrSize - 1; j >= index; j-- {
		l.arr[j+1] = l.arr[j]
	}
	l.arr[index] = num
	l.arrSize++
	return
}

func (l *myList) remove(index int) int {
	if index < 0 || index >= l.arrSize {
		panic("索引越界")
	}
	num := l.arr[index]

	for j := index; j < l.arrSize-1; j++ {
		l.arr[j] = l.arr[j+1]
	}
	l.arrSize--
	return num
}

func (l *myList) extendCapacity() {
	l.arr = append(l.arr, make([]int, l.arrCap*l.extendRatio)...)
	l.arrCap = len(l.arr)
	return
}

func (l *myList) toArray() []int {
	return l.arr[:l.arrSize]
}

func main() {
	nums := newMyList()

	nums.add(0)
	nums.add(1)
	nums.add(2)
	nums.add(3)
	nums.add(4)
	fmt.Println(nums)

	fmt.Println(nums.get(4))

	fmt.Println(nums.remove(4))
	fmt.Println(nums)
	nums.extendCapacity()
	fmt.Println(nums)
}

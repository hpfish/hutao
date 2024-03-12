package main

type node struct {
	val  int
	next *node
}

func newNode(val int) *node {
	return &node{val: val}
}

func function() int {
	return 0
}

func algorithm(n int) int { // 输入数据
	const a = 0      // 暂存数据（常量）
	b := 0           // 暂存数据（变量）
	newNode(0)       // 暂存数据（对象）
	c := function()  // 栈帧空间（调用函数）
	return a + b + c // 输出数据
}

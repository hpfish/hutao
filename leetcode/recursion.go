package main

import "container/list"

func tailRecur(n int, res int) int {
	if n == 0 {
		return res
	}
	return tailRecur(n-1, res+n)
}

// 使用迭代模式模拟递归
func forLoopRecur(n int) int {
	stack := list.New()
	res := 0
	// 递： 递归调用
	for i := n; i > 0; i-- {
		// 通过“入栈操作”模拟“递”
		stack.PushBack(i)
	}
	// 归： 返回结果
	for stack.Len() != 0 {
		// 通过“出栈操作”模拟归
		res += stack.Back().Value.(int)
		stack.Remove(stack.Back())
	}
	return res
}

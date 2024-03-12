package main

import "container/list"

// 层序遍历
func levelOrder(root *TreeNode) []any {
	// 初始化队列，加入根节点
	queue := list.New()
	queue.PushBack(root)

	nums := make([]any, 0)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		nums = append(nums, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return nums
}

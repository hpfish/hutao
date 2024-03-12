package main

var (
	nums []int
)

// 前序遍历
func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	nums = append(nums, node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

// 中序遍历
func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	nums = append(nums, node.Val)
	inOrder(node.Right)
}

// 后序遍历
func postOrder(node *TreeNode) {
	if node == nil {
		return
	}
	postOrder(node.Left)
	postOrder(node.Right)
	nums = append(nums, node.Val)
}

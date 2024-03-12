package main

type binarySearchTree struct {
	root *TreeNode
}

func newBinarySearchTree() *binarySearchTree {
	bst := &binarySearchTree{}
	bst.root = nil
	return bst
}

func (bst *binarySearchTree) getRoot() *TreeNode {
	return bst.root
}

func (bst *binarySearchTree) search(num int) *TreeNode {
	node := bst.root
	for node != nil {
		if node.Val < num {
			node = node.Right
		} else if node.Val > num {
			node = node.Left
		} else {
			break
		}
	}
	return node
}

func (bst *binarySearchTree) inert(num int) {
	cur := bst.root
	if cur == nil {
		bst.root = NewTreeNode(num)
		return
	}

	var pre *TreeNode = nil
	for cur != nil {
		if cur.Val == num {
			return
		}
		pre = cur
		if cur.Val < num {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}

	node := NewTreeNode(num)
	if pre.Val < num {
		pre.Right = node
	} else {
		pre.Left = node
	}
}

func (bst binarySearchTree) remove(num int) {
	cur := bst.root
	// 若数为空， 直接提前返回
	if cur == nil {
		return
	}

	var pre *TreeNode = nil
	for cur != nil {
		if cur.Val == num {
			break
		}
		pre = cur
		if cur.Val < num {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	if cur == nil {
		return
	}

	if cur.Left == nil || cur.Right == nil {
		var child *TreeNode = nil
		if cur.Left != nil {
			child = cur.Left
		} else {
			child = cur.Right
		}

		if cur != bst.root {
			if pre.Left == cur {
				pre.Left = child
			} else {
				pre.Right = child
			}
		} else {
			bst.root = child
		}
	} else {
		tmp := cur.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		bst.remove(tmp.Val)
		cur.Val = tmp.Val
	}
}

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	preHead := &ListNode{
		Next: head,
	}
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	reNum := n
	node := reK(preHead, k, reNum)
	for node != nil {
		reNum -= k
		node = reK(node, k, reNum)
	}
	return preHead.Next
}

func reK(preHead *ListNode, k int, reNodeNum int) *ListNode {
	end := preHead.Next

	if reNodeNum < k {
		return nil
	}

	i := 0
	for node := preHead.Next; i < k-1; i++ {
		nextNode := node.Next
		node.Next = nextNode.Next
		nextNode.Next = preHead.Next
		preHead.Next = nextNode
	}

	return end
}

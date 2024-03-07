package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func insertNode(n0 *ListNode, P *ListNode) {
	n1 := n0.Next
	P.Next = n1
	n0.Next = P
}

func removeItem(n0 *ListNode) {
	if n0.Next == nil {
		return
	}
	P := n0.Next
	n1 := P.Next
	n0.Next = n1
}

func access(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	return head
}

func findNode(head *ListNode, target int) int {
	index := 0
	for head != nil {
		if head.Val == target {
			return index
		}
		head = head.Next
		index++
	}
	return -1
}

type DoublyListNode struct {
	Val  int
	Next *DoublyListNode
	Prev *DoublyListNode
}

func NewDoublyListNode(val int) *DoublyListNode {
	return &DoublyListNode{
		Val:  val,
		Next: nil,
		Prev: nil,
	}
}

func main() {
	n0 := NewListNode(1)
	n1 := NewListNode(2)
	n2 := NewListNode(3)
	n3 := NewListNode(4)
	n4 := NewListNode(5)

	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	fmt.Println(n0)
}

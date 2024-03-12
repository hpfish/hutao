package main

import "testing"

func TestReverseKGroup(t *testing.T) {
	five := &ListNode{
		Val:  5,
		Next: nil,
	}

	four := &ListNode{
		Val:  4,
		Next: five,
	}

	three := &ListNode{
		Val:  3,
		Next: four,
	}
	second := &ListNode{
		Val:  2,
		Next: three,
	}
	head := &ListNode{
		Val:  1,
		Next: second,
	}

	result := reverseKGroup(head, 3)
	if result.Val != 3 {
		t.Fatalf("invalid value %d", result.Val)
	}
}

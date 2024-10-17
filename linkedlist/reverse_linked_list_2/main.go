package main

import (
	"fmt"
)

/*
Given the head of a singly linked list and two integers left and right where left <= right,
reverse the nodes of the list from position left to position right, and return the reversed list.
https://www.youtube.com/watch?v=RF_M9tX4Eag
*/
func main() {
	n5 := &ListNode{Val: 5, Next: nil}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}

	left := 3
	right := 4

	head := reverseBetween(n1, left, right)

	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Next: head}
	current := dummy

	// Move `current` to the node before the `left`th node
	for i := 0; i < left-1; i++ {
		current = current.Next
	}

	preLeft := current
	current = current.Next
	var prev *ListNode = nil

	for i := 0; i < right-left+1; i++ {
		tmp := current.Next
		current.Next = prev
		prev = current
		current = tmp
	}
	preLeft.Next.Next = current
	preLeft.Next = prev

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

package main

/*
The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100
*/
func main() {
	n3 := &ListNode{Val: 1, Next: nil}
	n2 := &ListNode{Val: 1, Next: n3}
	n1 := &ListNode{Val: 2, Next: n2}

	head := deleteDuplicates(n1)

	for head != nil {
		print(head.Val, " ")
		head = head.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

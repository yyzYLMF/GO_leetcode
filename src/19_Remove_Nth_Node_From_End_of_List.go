package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	var p, q *ListNode
	p, q = head, head
	for {
		if n <= 0 {
			break
		}
		p = p.Next
		n--
	}
	if p == nil {
		return head.Next
	}
	for {
		if p.Next == nil {
			break
		}
		p = p.Next
		q = q.Next
	}
	q.Next = q.Next.Next

	return head
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	fmt.Println(removeNthFromEnd(head, 1))
	return
}

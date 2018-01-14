package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tmp *ListNode
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	head, tmp = nil, nil
	p, q := l1, l2
	for {
		if p == nil || q == nil {
			break
		}
		if tmp == nil {
			if p.Val <= q.Val {
				head = p
				p = p.Next
			} else {
				head = q
				q = q.Next
			}
			tmp = head
		} else {
			if p.Val <= q.Val {
				tmp.Next = p
				p = p.Next
			} else {
				tmp.Next = q
				q = q.Next
			}
			tmp = tmp.Next
		}
	}
	if p == nil {
		tmp.Next = q
	} else {
		tmp.Next = p
	}
	return head
}

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	fmt.Println(mergeTwoLists(l1, l2))
	return
}

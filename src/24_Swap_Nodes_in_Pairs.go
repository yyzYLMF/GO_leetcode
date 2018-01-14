package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var ret, tmp *ListNode = nil, nil
	p, q := head, head.Next
	for {
		if ret == nil {
			ret, tmp = q, q
		} else {
			tmp.Next = q
		}
		p.Next = q.Next
		q.Next = p
		tmp = p
		p = p.Next
		if p != nil && p.Next != nil {
			q = p.Next
		} else {
			tmp.Next = p
			break
		}
	}
	return ret
}

func swapPairs2(head *ListNode) *ListNode {
	var odd, even *ListNode
	if head == nil || head.Next == nil {
		return head
	}
	odd, even = head, head.Next
	p, q, next := odd, even, head.Next.Next
	count := 1
	for {
		if next == nil {
			break
		}
		if count%2 == 1 {
			p.Next = next
			p = p.Next
			p.Next = nil
		} else {
			q.Next = next
			q = q.Next
			q.Next = nil
		}
		next = next.Next
		count++
	}
	nhead := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp := nhead
	count = 0
	for {
		if odd == nil || even == nil {
			break
		}
		if count%2 == 1 {
			tmp.Next = odd
			odd = odd.Next
		} else {
			tmp.Next = even
			even = even.Next
		}
		tmp = tmp.Next
		tmp.Next = nil
		count++
	}
	tmp.Next = odd
	return nhead.Next
}

func main() {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	n := swapPairs(input)
	fmt.Println(n.Val, n.Next.Val)
	return
}

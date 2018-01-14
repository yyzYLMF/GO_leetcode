package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l1 == nil {
		return l2
	}
	c := 0
	var ret, tmp *ListNode
	p, q := l1, l2
	for {
		if p == nil || q == nil {
			break
		}
		if ret == nil {
			ret = &ListNode{
				Val:  (p.Val + q.Val + c) % 10,
				Next: nil,
			}
			tmp = ret
		} else {
			tmp.Next = &ListNode{
				Val:  (p.Val + q.Val + c) % 10,
				Next: nil,
			}
			tmp = tmp.Next
		}
		c = (p.Val + q.Val + c) / 10
		p = p.Next
		q = q.Next
	}
	var left *ListNode
	if p != nil {
		left = p
	} else {
		left = q
	}
	for {
		if left == nil {
			break
		}
		tmp.Next = &ListNode{
			Val:  (left.Val + c) % 10,
			Next: nil,
		}
		tmp = tmp.Next

		c = (left.Val + c) / 10
		left = left.Next
	}
	if c != 0 {
		tmp.Next = &ListNode{
			Val:  c,
			Next: nil,
		}
	}
	return ret
}

func main() {
	first := &ListNode{
		Val:  1,
		Next: nil,
	}
	first.Next = &ListNode{
		Val:  8,
		Next: nil,
	}
	second := &ListNode{
		Val:  0,
		Next: nil,
	}
	ret := addTwoNumbers(first, second)
	for {
		if ret == nil {
			break
		}
		fmt.Println(ret.Val)
		ret = ret.Next
	}
	return
}

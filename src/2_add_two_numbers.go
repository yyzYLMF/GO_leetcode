package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	p := l1
	q := l2
	c := 0
	tmp := 0
	var head, tail *ListNode
	head = nil
	for {
		if p == nil || q == nil {
			break
		}
		tmp = (p.Val + q.Val + c) % 10
		c = (p.Val + q.Val + c) / 10
		var tmpNode = new(ListNode)
		tmpNode.Val = tmp
		tmpNode.Next = nil
		if head == nil {
			head = tmpNode
			tail = tmpNode
		} else {
			tail.Next = tmpNode
			tail = tail.Next
		}
		p = p.Next
		q = q.Next
	}
	if p == nil {
		p = q
	}
	for {
		if p == nil {
			break
		}
		tmp = (p.Val + c) % 10
		c = (p.Val + c) / 10
		var tmpNode = new(ListNode)
		tmpNode.Val = tmp
		tmpNode.Next = nil
		tail.Next = tmpNode
		tail = tail.Next
		p = p.Next
	}
	if c != 0 {
		var tmpNode = new(ListNode)
		tmpNode.Val = c
		tmpNode.Next = nil
		tail.Next = tmpNode
		tail = tail.Next
	}
	return head
}

func main() {
	fmt.Println("succ")
	return
}

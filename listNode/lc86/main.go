package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Add(val int) *ListNode {
	n.Next = &ListNode{Val: val}
	return n.Next
}

func (n *ListNode) Print() {
	c := n
	for c != nil {
		fmt.Println(c.Val)
		c = c.Next
	}
}

func main() {
	a := &ListNode{Val: 1}
	a.Add(4).Add(3).Add(2).Add(5).Add(2)
	r := partition(a, 3)
	r.Print()
}

func partition(head *ListNode, x int) *ListNode {
	smallDummy := &ListNode{}
	small := smallDummy
	bigDummy := &ListNode{}
	big := bigDummy
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			big.Next = head
			big = big.Next
		}
		head = head.Next
	}
	big.Next = nil
	small.Next = bigDummy.Next
	return smallDummy.Next
}

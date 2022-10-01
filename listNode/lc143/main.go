package main

import "fmt"

func main() {
	n := &ListNode{Val: 1}
	n.Add(2).Add(3).Add(4)
	reorderList(n)
}

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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	m := mid(head)
	middle := m.Next
	m.Next = nil
	reorder := reverse(middle)
	merge(head, reorder)
}

func mid(head *ListNode) *ListNode {
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	flag := false
	for l1 != nil && l2 != nil {
		if flag {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
		flag = !flag
	}
	return dummy.Next
}

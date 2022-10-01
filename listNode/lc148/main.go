package main

import "fmt"

func main() {
	l := &ListNode{Val: 4}
	l.Add(2).Add(1).Add(3)
	list := sortList(l)
	list.Print()
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

func sortList(head *ListNode) *ListNode {

	return mergeSort(head)
}

func findMid(head *ListNode) *ListNode {
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func mergeSort(head *ListNode) *ListNode {
	// 如果只有一个节点 直接就返回这个节点
	if head == nil || head.Next == nil {
		return head
	}
	// find middle
	middle := findMid(head)
	// 断开中间节点
	tail := middle.Next
	middle.Next = nil
	left := mergeSort(head)
	right := mergeSort(tail)
	result := mergeList(left, right)
	return result
}

func mergeList(l1, l2 *ListNode) *ListNode {
	headDummy := &ListNode{}
	head := headDummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	for l1 != nil {
		head.Next = l1
		l1 = l1.Next
		head = head.Next
	}
	for l2 != nil {
		head.Next = l2
		l2 = l2.Next
		head = head.Next
	}
	return headDummy.Next
}

package main

import "fmt"

/*
剑指 Offer 52. 两个链表的第一个公共节点
*/
func main() {

	p := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
	}
	fmt.Println(getIntersectionNode(&ListNode{
		Val:  1,
		Next: p,
	}, &ListNode{
		Val: 5,
		Next: &ListNode{
			Val:  0,
			Next: p,
		}}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

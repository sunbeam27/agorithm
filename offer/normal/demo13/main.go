package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(reverseList(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}))
}

func reverseList(head *ListNode) *ListNode {
	pre := head
	cur := head
	for cur != nil {
		next := cur.Next
		next.Next = pre
		pre = next
		cur = cur.Next
	}
	return pre
}

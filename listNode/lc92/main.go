package main

import "fmt"

func main() {
	head := new(ListNode)
	head.Val = 1
	head.Next = &ListNode{
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
	}
	node := reverseBetween(head, 2, 4)

	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := new(ListNode)
	// 虚拟头指针
	dummy.Next = head

	pre := dummy
	// 找到左边的前一个
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	rightNode := pre
	// 右边的节点
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	start := pre.Next

	cur := rightNode.Next

	pre.Next = nil
	rightNode.Next = nil

	reverse(start)

	pre.Next = rightNode
	start.Next = cur

	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

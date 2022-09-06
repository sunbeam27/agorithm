package main

func main() {
	isPalindrome(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	})
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	stack := make([]int, 0)
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	n := len(stack)
	for i, v := range stack[:n/2] {
		if v != stack[n-1-i] {
			return false
		}
	}
	return true
}

package main

func main() {
	addTwoNumbers(&ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}, &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	})
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	val, carry := 0, 0
	r := &ListNode{}
	res := r
	for l1 != nil || l2 != nil {
		n1 := 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		n2 := 0
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		val, carry = sum%10, sum/10
		r.Next = &ListNode{
			Val: val,
		}
		r = r.Next
	}
	if carry != 0 {
		r.Next = &ListNode{
			Val: carry,
		}
	}
	return res.Next
}

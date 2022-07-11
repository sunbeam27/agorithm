package main

import "fmt"

func main() {
	fmt.Println(levelOrder(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{{root.Val}}
	res = nextFloor(root.Left, root.Right, res, 1)
	return res
}

func nextFloor(left, right *TreeNode, res [][]int, floor int) [][]int {
	if left == nil && right == nil {
		return res
	}

	if left != nil {
		if floor > len(res)-1 {
			res = append(res, []int{})
		}
		res[floor] = append(res[floor], left.Val)
		res = nextFloor(left.Left, left.Right, res, floor+1)
	}
	if right != nil {
		if floor > len(res)-1 {
			res = append(res, []int{})
		}
		res[floor] = append(res[floor], right.Val)
		res = nextFloor(right.Left, right.Right, res, floor+1)
	}

	return res
}

package main

import "fmt"

func main() {
	fmt.Println(maxDepth(&TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}))
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var l, r = 1, 1

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dfs(root)
	if l > r {
		return l
	}
	return r
}

func dfs(root *TreeNode) {
	if root.Right != nil {
		dfs(root.Right)
		r++
	}
	if root.Left != nil {
		dfs(root.Left)
		l++
	}
}

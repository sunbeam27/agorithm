package main

import "fmt"

// 两数之和
func main() {
	fmt.Println(twoSum([]int{1, 2, 4, 6, 10}, 3))
}
func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] > target {
			j--
			continue
		}
		if numbers[i]+numbers[j] == target {
			return []int{i, j}
		}
		if numbers[i]+numbers[j] < target {
			i++
			continue
		}
	}
	return nil
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) || len(preorder) == 0 {
		return nil
	}
	head := &TreeNode{preorder[0], nil, nil}
	idx := 0
	for ; idx < len(inorder); idx++ {
		if inorder[idx] == preorder[0] {
			break
		}
	}
	head.Left = buildTree(preorder[1:len(inorder[:idx])+1], inorder[:idx])
	head.Right = buildTree(preorder[len(inorder[:idx])+1:], inorder[idx+1:])
	return head
}

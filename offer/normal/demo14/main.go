package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return isEqual(A.Left, B) || isEqual(A.Right, B)
}

func isEqual(A *TreeNode, B *TreeNode) bool {
	if A == nil {
		return false
	}
	if B == nil {
		return true
	}
	if A.Val != B.Val {
		return false
	}
	return isEqual(A.Left, B.Left) && isEqual(A.Right, B.Right)
}

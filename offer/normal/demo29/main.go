package main

func main() {

}

func findTarget(root *TreeNode, k int) bool {
	m := map[int]struct{}{}
	var f func(*TreeNode)
	f = func(node *TreeNode) {
		if node != nil {
			f(node.Left)
			m[node.Val] = struct{}{}
			f(node.Right)
		}
	}
	f(root)
	for val := range m {
		_, ok := m[k-val]
		if ok {
			return true
		}
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

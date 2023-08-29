package main

import "fmt"

type Tree struct {
	val         int
	left, right *Tree
}

func main() {
	root := &Tree{
		val:   1,
		left:  &Tree{val: 2},
		right: &Tree{val: 3},
	}
	fmt.Println(mid(root))
}

func mid(root *Tree) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	var st []*Tree
	for root != nil || len(st) != 0 {
		if root != nil {
			st = append(st, root)
			root = root.left
		} else {
			root = st[len(st)-1]
			res = append(res, root.val)
			st = st[:len(st)-1]
			root = root.right
		}
	}
	return res
}

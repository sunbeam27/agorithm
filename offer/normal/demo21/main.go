package main

import "fmt"

// 中心数
func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}

func pivotIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		left, right := 0, 0
		for j := 0; j <= i; j++ {
			left += nums[j]
		}
		for k := i; k < len(nums); k++ {
			right += nums[k]
		}
		if left == right {
			return i
		}
	}
	return -1
}

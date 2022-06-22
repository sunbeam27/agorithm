package main

import "fmt"

/*
	轮转数组
	输入 nums = [1,2,3,4,5,6,7] , k=3
	输出 [5,6,7,1,2,3,4]
*/
func main() {
	fmt.Println(rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3))
}
func rotate(nums []int, k int) []int {
	k %= len(nums)
	return append(nums[len(nums)-k:], nums[:len(nums)-k]...)
}

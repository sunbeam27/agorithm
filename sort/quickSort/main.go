package main

import "fmt"

func main() {
	nums := []int{5, 2, 3, 1}
	//arr := quick(nums)
	//fmt.Println(arr)
	//sort.Slice(nums, func(i, j int) bool {
	//	return nums[i] < nums[j]
	//})
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

// 快速排序入口函数
func quickSort(nums []int, l int, r int) {
	// 递归终止条件
	if l >= r {
		return
	}
	p := nums[l]
	i, j := l, r
	for i < j {
		for i < j && p >= nums[j] {
			j--
		}
		for i < j && p <= nums[i] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[r] = nums[r], nums[i]
	quickSort(nums, l, i-1)
	quickSort(nums, i+1, r)
}

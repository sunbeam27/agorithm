package main

import "fmt"

func main() {
	fmt.Println(mergeSort([]int{1, 5, 435, 3, 2, 6435, 7, 47}))
}

func mergeSort(nums []int) []int {
	return merge1(nums)
}

func merge1(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := merge1(nums[:mid])
	right := merge1(nums[mid:])
	return merge(left, right)
}

func merge(n1, n2 []int) []int {
	res := make([]int, 0, len(n1)+len(n2))
	i1, i2 := 0, 0
	for i1 < len(n1) && i2 < len(n2) {
		if n1[i1] < n2[i2] {
			res = append(res, n1[i1])
			i1++
		} else {
			res = append(res, n2[i2])
			i2++
		}
	}
	res = append(res, n1[i1:]...)
	res = append(res, n2[i2:]...)

	return res
}

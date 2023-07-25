package main

import "fmt"

func main() {
	fmt.Println(mergeSort([]int{1, 5, 435, 3, 2, 6435, 7, 47}))
}

func mergeSort(nums []int) []int {
	return mergesort(nums)
}
func mergesort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergesort(arr[:mid])
	right := mergesort(arr[mid:])
	return merge(left, right)
}

func merge(l1, l2 []int) []int {
	t := make([]int, 0, len(l1)+len(l2))
	i1, i2 := 0, 0
	for i1 < len(l1) && i2 < len(l2) {
		if l1[i1] < l2[i1] {
			t = append(t, l1[i1])
			i1++
		} else {
			t = append(t, l2[i2])
			i2++
		}
	}
	t = append(t, l1[i1:]...)
	t = append(t, l2[i2:]...)

	return t
}

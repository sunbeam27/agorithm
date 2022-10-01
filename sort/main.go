package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 1, 4, 7, 5, 3}
	//quickSort(arr, 0, len(arr)-1)
	//bubble(arr)
	quick(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func bubbleSort(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	return arr
}

// 快速排序入口函数
func quickSort(nums []int, left, right int) {
	// 递归终止条件
	if left >= right {
		return
	}
	p := partition(nums, left, right)
	quickSort(nums, left, p-1)
	quickSort(nums, p+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[right], nums[i] = nums[i], pivot
	return i
}

func bubble(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func quick(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := nums[left]
	i := right
	for j := right; j > left; j-- {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i--
		}
	}
	nums[right], nums[i] = nums[i], pivot
	quick(nums, left, i-1)
	quick(nums, i+1, right)
}

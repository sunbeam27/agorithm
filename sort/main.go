package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 1, 4, 7, 5, 3}
	//quickSort(arr, 0, len(arr)-1)
	//bubbleSort(arr)
	quick(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
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

func quick(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := arr[left]
	i, j := left, right
	for i < j {
		for i < j && pivot <= arr[j] {
			j--
		}
		for i < j && pivot >= arr[i] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[left] = arr[left], arr[i]
	quick(arr, left, i-1)
	quick(arr, i+1, right)
}

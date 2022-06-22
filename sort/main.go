package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{1, 2, 4, 7, 5, 4, 3}
	sort := selectionSort(ints)
	fmt.Println(sort)
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

// todo!!!!!
func quickSort(arr []int) []int {
}

func qsourt(arr []int, low, high int) {

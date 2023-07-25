package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 7, 8, 9, 11}, 4))
}

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := (l + r) / 2
		if arr[m] == target {
			return m
		} else if arr[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

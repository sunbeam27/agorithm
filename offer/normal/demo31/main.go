package main

import (
	"fmt"
)

func main() {
	fmt.Println(relativeSortArray([]int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6}))
}
func relativeSortArray(arr1 []int, arr2 []int) []int {
	m2 := make(map[int]int, len(arr2))
	for _, v := range arr2 {
		m2[v] = 0
	}
	res2 := make([]int, 0)
	res1 := make([]int, 0)
	for j := 0; j < len(arr1); j++ {
		if _, ok := m2[arr1[j]]; ok {
			m2[arr1[j]] += 1
		} else {
			res2 = append(res2, arr1[j])
		}
	}
	for _, v := range arr2 {
		for i := 0; i < m2[v]; i++ {
			res1 = append(res1, v)
		}
	}
	return append(res1, res2...)
}

package main

func quick(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	p := arr[0]
	low := make([]int, 0)
	high := make([]int, 0)
	equal := make([]int, 0)
	for _, v := range arr {
		if v > p {
			high = append(high, v)
		} else if v < p {
			low = append(low, v)
		} else {
			equal = append(equal, v)
		}
	}
	l, h := quick(low), quick(high)
	return append(append(l, equal...), h...)
}

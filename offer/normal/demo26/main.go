package main

import "fmt"

func main() {
	fmt.Println(lastRemaining(5, 3))
}

func lastRemaining(n int, m int) int {
	idx := 0
	for i := 2; i <= n; i++ {
		idx = (idx + m) % i
	}
	return idx
}

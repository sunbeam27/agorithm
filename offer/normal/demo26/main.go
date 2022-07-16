package main

import "fmt"

func main() {
	fmt.Println(lastRemaining(5, 3))
}

func lastRemaining(n int, m int) int {
	circle := make([]int, n)
	for i := 0; i < len(circle); i++ {
		circle[i] = i
	}
	for len(circle) > 1 {
		circle = append(circle[:m%len(circle)], circle[m%len(circle)+1:]...)
	}
	return circle[0]
}

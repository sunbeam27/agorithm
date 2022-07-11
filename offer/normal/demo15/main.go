package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}}))
}
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	data := make([]int, 0, len(matrix)*len(matrix[0]))
	a, b, x, y := len(matrix), len(matrix[0]), 0, 0
	for a != 0 && b != 0 {
		for i := x; i < b-1; i++ {
			data = append(data, matrix[x][i])
		}
		for i := y; i < a-1; i++ {
			data = append(data, matrix[i][b-1])
		}
		for i := b - 1; i >= x; i-- {
			data = append(data, matrix[a-1][i])
		}
		for i := a - 1; i >= y; i-- {
			data = append(data, matrix[i][x])
		}
	}
	return data
}

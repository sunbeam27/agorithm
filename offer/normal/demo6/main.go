package main

import "fmt"

/*
给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。


例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。




示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false

*/
func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if search(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, i, j, k int, word string) bool {
	if k == len(word) {
		return true
	}
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return false
	}
	if word[k] == board[i][j] {
		tmp := board[i][j]
		board[i][j] = ' '
		if search(board, i+1, j, k+1, word) ||
			search(board, i, j+1, k+1, word) ||
			search(board, i-1, j, k+1, word) ||
			search(board, i, j-1, k+1, word) {
			return true
		} else {
			board[i][j] = tmp
		}
	}
	return false
}

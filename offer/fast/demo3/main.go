package main

import (
	"fmt"
	"strconv"
)

/*
给定一个非负整数 n，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。


示例 1:

输入: n = 2
输出: [0,1,1]
解释:
0 --> 0
1 --> 1
2 --> 10


*/
func main() {
	fmt.Println(countBits(2))
}

func countBits(n int) []int {
	res := make([]int, 0, n)
	for i := 0; i <= n; i++ {
		fInt := strconv.FormatInt(int64(i), 2)
		count := 0
		for _, c := range fInt {
			if c == '1' {
				count++
			}
		}
		res = append(res, count)
	}
	return res
}

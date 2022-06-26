package main

import "fmt"

/*
实现pow(x,n)，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。


示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000

*/
func main() {
	fmt.Println(myPow(2.00000, -10))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	t := myPow(x, n/2)
	if n%2 == 0 {
		return t * t
	}
	return x * t * t
}

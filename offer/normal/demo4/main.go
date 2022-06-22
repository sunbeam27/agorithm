package main

import "fmt"

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n级的台阶总共有多少种跳法。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

*/
func main() {
	fmt.Println(numWays(3))
}

func numWays(n int) int {
	a := 1
	b := 1
	for i := 0; i < n-1; i++ {
		b, a = (a+b)%1000000007, b
	}
	return b
}

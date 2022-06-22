package main

import "fmt"

func main() {
	fmt.Println(divide(15, 2))

}

/*
给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%' 。



注意：

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231, 231−1]。本题中，如果除法结果溢出，则返回 231 − 1

*/
func divide(a int, b int) int {
	s := 0
	if a == 0 {
		return 0
	}
	fd := 0
	if a < 0 {
		a = -a
		fd += 1
	}
	if b < 0 {
		b = -b
		fd += 1
	}
	for {
		a -= b
		if a <= 0 {
			if fd == -1 {
				return -s
			}
			return s
		}
		s += 1
	}
}

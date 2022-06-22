package main

import (
	"fmt"
	"strconv"
)

/*
给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。

输入为 非空 字符串且只包含数字 1 和 0。
*/
func main() {

	fmt.Println(addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
		"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
}

func addBinary(a string, b string) string {
	parseA, _ := strconv.ParseInt(a, 2, 64)
	parseB, _ := strconv.ParseInt(b, 2, 64)
	return strconv.FormatInt(parseA+parseB, 2)
}

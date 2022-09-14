package main

import (
	"fmt"
	"time"
)

/*
给你一个字符串 s，找到 s 中最长的回文子串。



示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"


*/
func main() {
	parse, err := time.Parse("2006-01-02", "2022-09-08")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parse.Weekday())
	fmt.Println(longestPalindrome("aa"))
}

func longestPalindrome(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}
	max := ""
	for i := 0; i < n; i++ {

		for j := i + 1; j <= n; j++ {
			if isReverse(s[i:j]) && j-i > len(max) {
				max = s[i:j]
			}
		}
	}
	return max
}

func isReverse(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
			continue
		} else {
			return false
		}
	}

	return true

}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("  hello world!  "))
}

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	words := strings.Fields(s)
	result := ""
	for i := len(words) - 1; i >= 0; i-- {
		result += words[i]
		if i != 0 {
			result += " "
		}
	}
	return result

}

package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("0P"))
}

func isPalindrome(s string) bool {
	if len(s) == 1 || len(s) == 0 {
		return true
	}
	s = lowCase(s)
	i := 0
	j := len(s) - 1
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
func lowCase(s string) string {
	toLow := 'a' - 'A'
	str := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if ('a' <= (s[i]) && (s[i]) <= 'z') || ('0' <= (s[i]) && (s[i]) <= '9') {
			str = append(str, s[i])
		} else if 'A' <= byte(s[i]) && byte(s[i]) <= 'Z' {
			str = append(str, byte(int32(s[i])+toLow))
		}
	}
	return string(str)
}

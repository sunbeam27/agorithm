package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("adc", "dcda"))
	fmt.Println(checkInclusion2("adc", "dcda"))
}

func checkInclusion(s1 string, s2 string) bool {
	m := make(map[byte]int, len(s1))
	for i := 0; i < len(s1); i++ {
		m[s1[i]]++
	}
	for l, r := 0, len(s1); r < len(s2)+1; r++ {
		t := s2[l:r]
		m1 := make(map[byte]int, len(t))
		for i := 0; i < len(t); i++ {
			m1[t[i]]++
		}
		flag := false
		for k := range m1 {
			if m1[k] != m[k] {
				flag = false
				break
			}
			flag = true
		}
		if flag {
			return true
		}
		l++
	}
	return false
}

func checkInclusion2(s1 string, s2 string) bool {

	m := [26]int{}
	for _, c := range s1 {
		m[c-'a']--
	}

	left := 0
	for i := 0; i < len(s2); i++ {
		m[s2[i]-'a']++
		if m[s2[i]-'a'] > 0 {
			m[s2[left]-'a']--
			left++
		}
		if i-left+1 == len(s1) {
			return true
		}
	}

	return false
}

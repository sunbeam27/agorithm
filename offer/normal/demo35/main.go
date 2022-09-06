package main

import "fmt"

func main() {
	fmt.Println(isAnagram("rat", "car"))
}

func isAnagram(s string, t string) bool {
	chars := [26]int{}
	for _, c := range s {
		chars[int(c-'a')] += 1
	}
	for _, c := range t {
		chars[int(c-'a')] -= 1
	}
	sum := 0
	for _, s := range chars {
		if s < 0 {
			s = -s
		}
		sum += s
	}
	return sum == 0
}

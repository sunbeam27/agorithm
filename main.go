package main

import (
	"container/list"
	"fmt"
	"strings"
)

type A struct {
	Name *string
	list.List
}

/*
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a  student. "，则输出"student. a am I"


*/
func main() {
	fmt.Println(Reverse("I am a  student. "))
}

func Reverse(str string) string {
	split := strings.Fields(str)
	res := ""
	for i := len(split) - 1; i >= 0; i-- {
		if i != len(split)-1 {
			res += split[i] + " "
		} else {
			res += split[i] + " "
		}
	}
	return res
}

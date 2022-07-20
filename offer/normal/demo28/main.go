package main

import "fmt"

func main() {
	fmt.Println(judge("abBAcd"))
	nums := &[]int{4, 5, 6, 1, 1, 2}
	move(*nums, 1)
	fmt.Println(nums)
}

//1.数组 和 target 把 target 移动到数组头，操作原数组，其他不变
func move(nums []int, target int) {
	for i, num := range nums {
		if target == num {
			copy(nums[1:i+1], nums[0:i])
			nums[0] = target

		}
	}
	fmt.Println(nums)
}

// 2.字符串 判断 字符串中 同时有大小写存在的 最大的
// 如 abBAcd 结果为B
func judge(str string) string {
	ch := [52]int{}
	for _, c := range str {
		if c < 'Z' {
			ch[int(c-'A')] = 1
		} else {
			ch[int(c-'a')+26] = 1
		}
	}
	max := 0
	for i := 0; i < 26; i++ {
		if ch[i] == 1 && ch[i+26] == 1 {
			if i > max {
				max = i
			}
		}
	}
	return string([]byte{uint8(max + 65)})
}

// 3. 二维数组，找到最小的下降路径

func find() {

}

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var ch chan int // nil channel
	go func() {
		ch <- 123 // 阻塞
		fmt.Println("阻塞了")
	}()
	go func() {
		time.Sleep(1 * time.Minute)
	}()
	fmt.Println(<-ch) // 阻塞
}

func 随机实验数据() {
	for i := 0; i < 10; i++ {
		fmt.Println("-------------start-" + strconv.Itoa(i) + "次--------------------")
		chars := make([]string, 0)
		for i := 'A'; i < 'M'; i++ {
			chars = append(chars, string(byte(i)))
		}
		fmt.Println(RandArr(chars, 4))
		nums := make([]string, 0)
		for i := 1; i <= 56; i++ {
			nums = append(nums, strconv.Itoa(i))
		}
		fmt.Println(RandArr(nums, 4))
		fmt.Println("---------------end----------------------")
	}
}
func RandArr(arr []string, num int) []string {
	res := make([]string, 0, num)
	for i := 0; i < num; i++ {
		n := rand.Intn(len(arr))
		r := arr[n]
		res = append(res, r)
		arr = append(arr[:n], arr[n+1:]...)
	}
	return res
}

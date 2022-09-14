package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数到30就输的游戏

var win = [...]int{29, 26, 23, 20, 17, 14, 11, 8, 5, 2}

// 获胜策略
func AI(p int) (result int) {
	for _, num := range win {
		if p+1 == num || p+2 == num {
			return num
		}
	}
	return rand2(p)
}

// 随机方法
func rand2(a int) int {
	return a + 1 + rand.Intn(2) // 1, 2
}

func game() {
	rand.Seed(time.Now().UnixNano())

	a := 0
	p := 0

	var ai []int
	var player []int
	for {
		// 玩家先走
		p = rand2(a)
		player = append(player, p)
		if p >= 30 {
			fmt.Println("ai赢了这个游戏")
			fmt.Println("player:", player)
			fmt.Println("ai:", ai)
			return
		}
		// ai 知道策略
		a = AI(p)
		ai = append(ai, a)
		if a >= 30 {
			fmt.Println("玩家赢了这个游戏")
			fmt.Println("player:", player)
			fmt.Println("ai:", ai)
			return
		}
	}
}

func main() {
	game()
}

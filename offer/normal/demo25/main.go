package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isStraight([]int{1, 2, 3, 4, 5}))
}

func isStraight(nums []int) bool {
	king := 0
	cards := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			king++
		} else {
			cards = append(cards, nums[i])
		}
	}
	sort.Ints(cards)
	cur := cards[0]
	for i := 1; i < len(cards); i++ {
		if cards[i]-1 != cur {
			king--
			if king < 0 {
				return false
			}
		}
		cur = cards[i]
	}
	return true
}

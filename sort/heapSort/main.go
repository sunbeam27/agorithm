package main

import (
	"fmt"
)

// 插入（上浮）
func upAdjust(array []int) {
	childIndex := len(array) - 1
	parentIndex := (childIndex - 1) / 2
	//保存插入叶子节点值，用于最后的赋值
	temp := array[childIndex]

	for childIndex > 0 && temp < array[parentIndex] {
		//单项赋值，不用交换
		array[childIndex] = array[parentIndex]
		childIndex = parentIndex
		parentIndex = (childIndex - 1) / 2
	}
	array[childIndex] = temp
}

// 删除（下沉）
func downAdjust(array []int, parentIndex, length int) {
	//保存父节点的值，用于最后的赋值
	temp := array[parentIndex]
	childIndex := 2*parentIndex + 1
	for childIndex < length {
		//如果存在右孩子，并且右孩子小于左孩子的值，定位到右孩子
		if childIndex+1 < length && array[childIndex+1] < array[childIndex] {
			childIndex++
		}
		//如果父节点小于任何一个孩子的值，直接退出
		if temp <= array[childIndex] {
			break
		}
		array[parentIndex] = array[childIndex]
		parentIndex = childIndex
		childIndex = 2*childIndex + 1
	}
	array[parentIndex] = temp
}

// 构建最小二叉堆
func buildHeap(array []int) {
	for i := (len(array) - 2) / 2; i >= 0; i-- {
		downAdjust(array, i, len(array))
	}
}

func main() {
	array := []int{1, 3, 2, 6, 5, 7, 8, 9, 10, 0}
	upAdjust(array)
	fmt.Println(array)

	array = []int{7, 1, 3, 10, 5, 2, 8, 9, 6}
	buildHeap(array)
	fmt.Println(array)
}

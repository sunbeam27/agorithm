package main

import "fmt"

func main() {
	/*
	   	["MinStack","push","push","push","min","pop","top","min"]
	      [[],[-2],[0],[-3],[],[],[],[]]
	*/
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.Min())
	minStack.Pop()
	minStack.Top()
	fmt.Println(minStack.Min())
}

type MinStack struct {
	Data   []int
	MinVal []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Data:   make([]int, 0),
		MinVal: nil,
	}
}

func (this *MinStack) Push(x int) {
	this.Data = append(this.Data, x)
	if len(this.MinVal) == 0 || this.MinVal[len(this.MinVal)-1] >= x {
		this.MinVal = append(this.MinVal, x)
	}
}

func (this *MinStack) Pop() {
	pop := this.Data[len(this.Data)-1]
	if pop <= this.MinVal[len(this.MinVal)-1] {
		this.MinVal = append(this.MinVal[:len(this.MinVal)-1])
	}
	this.Data = append(this.Data[:len(this.Data)-1])
}

func (this *MinStack) Top() int {
	return this.Data[len(this.Data)-1]
}

func (this *MinStack) Min() int {
	return this.MinVal[len(this.MinVal)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

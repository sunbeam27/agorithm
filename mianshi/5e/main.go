package main

import (
	"fmt"
	"sync"
	"time"
)

/*
*
package main

import (

	"fmt"
	"sync"
	"time"

)

	func main() {
	    wg := sync.WaitGroup{}
	    c := make(chan struct{})
	    for i := 0; i < 10; i++ {
	        wg.Add(1)
	        go func(num int, close <-chan struct{}) {
	            defer wg.Done()
	            <-close
	          fmt.Println(num)
	        }(i, c)
	    }

	    if WaitTimeout(&wg, time.Second*5) {
	        close(c)
	        fmt.Println("timeout exit")
	    }
	    time.Sleep(time.Second * 10)
	}

	func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	    // 要求手写代码
	    // 要求sync.WaitGroup支持timeout功能
	    // 如果timeout到了超时时间返回true
	    // 如果WaitGroup自然结束返回false
	}
*/
func main() {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(num)
		}(i, c)
	}

	if WaitTimeout(&wg, time.Second*5) {
		close(c)
		fmt.Println("timeout exit")
	}
	time.Sleep(time.Second * 10)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false
	end := make(chan bool)
	go func() {
		wg.Wait()
		end <- true
	}()
	timer := time.NewTimer(timeout)
	select {
	case <-timer.C:
		return true
	case <-end:
		return false
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var val []int

func kthSmallest(root *TreeNode, k int) int {
	val = make([]int, 0)
	middle(root)

	return val[k-1]
}

func middle(root *TreeNode) {
	if root == nil {
		return
	}
	middle(root.Left)
	val = append(val, root.Val)
	middle(root.Right)
}

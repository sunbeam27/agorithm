package main

import (
	"fmt"
	"sync"
)

// 输出1-100 奇数偶数 按顺序 2个goroutine
func main() {
	jCh := make(chan struct{}, 1)
	oCh := make(chan struct{}, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 100; i += 2 {
			<-oCh
			fmt.Println(i)
			jCh <- struct{}{}
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i < 100; i += 2 {
			<-jCh
			fmt.Println(i)
			oCh <- struct{}{}
		}
		wg.Done()
	}()
	oCh <- struct{}{}

	wg.Wait()
}

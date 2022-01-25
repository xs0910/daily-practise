package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 使用WaitGroup 和 channel 控制 协程个数
func TestGoroutine(t *testing.T) {
	var wg = sync.WaitGroup{}
	manyDataList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan bool, 3)
	for _, v := range manyDataList {
		wg.Add(1)
		go func(data int) {
			defer wg.Done()
			ch <- true
			fmt.Printf("go func: %d, time: %d\n", data, time.Now().Unix())
			time.Sleep(time.Second)
			<-ch
		}(v)
	}
	wg.Wait()
}

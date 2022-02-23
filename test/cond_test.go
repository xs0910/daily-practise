package test

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var status int64

func TestCond(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go listen(c)
	}
	time.Sleep(1 * time.Second)
	go broadcast(c)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func broadcast(c *sync.Cond) {
	c.L.Lock()
	defer c.L.Unlock()
	atomic.StoreInt64(&status, 1)
	// c.Signal()    // 唤醒队列前面的一个Goroutine
	c.Broadcast() // 唤醒队列中全部的Goroutine
	// 唤醒的顺序也是按照加入队列的先后顺序，先加入的会先被唤醒，而后加入的可能需要等待调度器的调度
}

func listen(c *sync.Cond) {
	c.L.Lock()
	defer c.L.Unlock()
	for atomic.LoadInt64(&status) != 1 {
		c.Wait()
	}
	fmt.Println("listen:", time.Now())
}

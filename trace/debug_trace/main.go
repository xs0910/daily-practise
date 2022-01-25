package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello scheduler")
	}

	// 运行命令 GODEBUG=schedtrace=1000 go run main.go
	// GODEBUG=gctrace=1 go run main.go
	// 或者
	// go build .
	// GODEBUG=schedtrace=1000 ./debug_trace.exe
	// GODEBUG=schedtrace=1000,scheddetail=1 ./debug_trace.exe  # 加上scheddetail=1 可以打印更详细的trace信息

	// schedtrace： 设置 schedtrace=X 参数可以使运行时在每 X 毫秒发出一行调度器的摘要信息到标准 err 输出中。
	// scheddetail：设置 schedtrace=X 和 scheddetail=1 可以使运行时在每 X 毫秒发出一次详细的多行信息，信息内容主要包括调度程序、处理器、OS 线程 和 Goroutine 的状态。

	// 自旋线程的这个说法，是因为 Go Scheduler 的设计者在考虑了 “OS 的资源利用率” 以及 “频繁的线程抢占给 OS 带来的负载” 之后，
	// 提出了 “Spinning Thread” 的概念。也就是当 “自旋线程” 没有找到可供其调度执行的 Goroutine 时，并不会销毁该线程 ，
	// 而是采取 “自旋” 的操作保存了下来。虽然看起来这是浪费了一些资源，但是考虑一下 syscall 的情景就可以知道，
	// 比起 “自旋”，线程间频繁的抢占以及频繁的创建和销毁操作可能带来的危害会更大。
}

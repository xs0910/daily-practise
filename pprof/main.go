package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func main() {
	go func() {
		ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()

	tick := time.Tick(time.Second / 100)
	var buf []byte
	for range tick {
		buf = append(buf, make([]byte, 1024*1024)...)
	}

	// 输入网址 ip:port/debug/pprof
	// block：goroutine的阻塞信息，本例就截取自一个goroutine阻塞的demo，但block为0，没掌握block的用法
	// goroutine：所有goroutine的信息，下面的full goroutine stack dump是输出所有goroutine的调用栈，是goroutine的debug=2，后面会详细介绍。
	// heap：堆内存的信息
	// mutex：锁的信息
	// threadcreate：线程信息

	// 当连接在服务器终端上的时候，是没有浏览器可以使用的，Go提供了命令行的方式，能够获取以上5类信息，这种方式用起来更方便。
	//# 下载cpu profile，默认从当前开始收集30s的cpu使用情况，需要等待30s
	//go tool pprof http://localhost:6060/debug/pprof/profile   # 30-second CPU profile
	//go tool pprof http://localhost:6060/debug/pprof/profile?seconds=120     # wait 120s

	//# 下载heap profile
	//go tool pprof http://localhost:6060/debug/pprof/heap      # heap profile

	//# 下载goroutine profile
	//go tool pprof http://localhost:6060/debug/pprof/goroutine # goroutine profile

	//# 下载block profile
	//go tool pprof http://localhost:6060/debug/pprof/block     # goroutine blocking profile

	//# 下载mutex profile
	//go tool pprof http://localhost:6060/debug/pprof/mutex
}

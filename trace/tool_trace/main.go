package main

import (
	_ "net/http/pprof"
	"os"
	"runtime/trace"
)

// trace记录了运行时的信息，能提供可视化的Web页面。
func main() {
	// 创建trace文件
	f, err := os.Create("./trace/tool_trace/trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "Go 语言编程之旅"
	}()

	<-ch

	// 使用命令：go tool trace trace.out 执行可视化
}

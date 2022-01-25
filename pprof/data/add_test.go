package main

import "testing"

func TestAdd(t *testing.T) {
	_ = Add("go-programing-tour-book")
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add("go-programing-tour-book")
	}

	// git bash 终端执行命令
	// go test -bench=. -cpuprofile=cpu.profile
	// go test -bench=. -memprofile=mem.profile

	// 启动pprof 可视化界面
	// go tool pprof -http=:8080 cpu.profile
	// 或者
	// go tool pprof cpu.profile
	// (pprof) web

	// 安装pprof go get -u github.com/google/toppprof
	// 启动：pprof -http=:8080 cpu.prof 火焰图
}

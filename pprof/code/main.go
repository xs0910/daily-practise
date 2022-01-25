package main

import (
	"os"
	"runtime/pprof"
)

// 使用代码生成pprof文件
func main() {
	cpuOut, _ := os.Create("./pprof/code/cpu.out")
	defer cpuOut.Close()

	pprof.StartCPUProfile(cpuOut)
	defer pprof.StopCPUProfile()

	menOut, _ := os.Create("./pprof/code/mem.out")
	defer menOut.Close()
	defer pprof.WriteHeapProfile(menOut)

	Sum(3, 5)
}

func Sum(a, b int) int {
	return a + b
}

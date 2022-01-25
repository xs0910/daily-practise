package test

import (
	"log"
	"testing"
)

func TestRecoverAndPanic(t *testing.T) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("recover: %v", err)
			}
		}()

		panic("手动触发panic")
	}()

	log.Println("Go 语言编程之旅：一起用 Go 做项目")
	// 如果这里使用 t.Log, defer 的打印将不会出来
}

/*
让 Go Panic 的十种方法
1. 数组/切片越界
2. 空指针调用
3. 过早关闭HTTP响应体
4. 除以零
5. 向已关闭的通道发消息
6. 重复关闭通道
7. 关闭未初始化的通道
8. 未初始化map
9. 跨协程的恐慌处理
10.sync计数为负值
*/

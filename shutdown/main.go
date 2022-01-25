package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// 优雅关停例子
func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second * 5)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 将服务在goroutine 中启动
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit // 阻塞等待接收 channel 数据
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) // 5s 缓冲时间处理已有请求
	defer cancel()

	// 调用net/http 包提供的优雅关闭函数 Shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

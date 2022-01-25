package main

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	info()
	withField(uuid.New().String(), "127.0.0.1")
}

func info() {
	// 默认情况下使用的是TextFormatter，一经设置，就是全局的
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		DataKey:         "data", // 使用DataKey来包装自定义要输出的字段
		PrettyPrint:     true,   // 格式化输出
	})
	log.Info("A walrus appears")
}

func withField(requestId string, userIp string) {
	requestLogger := log.WithFields(log.Fields{"request_id": requestId, "user_id": userIp})
	requestLogger.Info("something happened on that request")
	requestLogger.Warn("something not great happened")
}

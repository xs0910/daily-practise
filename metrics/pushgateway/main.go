package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	// 统计请求数量
	httpRequestCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "service",
			Name:      "http_request_total",
			Help:      "Total number of http_request",
		},
	)

	// NewCounter 与 NewCounterVec 的区别，加入了label
	httpRequestCountVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "service",
			Name:      "http_request_total",
			Help:      "Total number of http_request",
		},
		[]string{"kind"},
	)

	// 监控实时并发量（处理中的请求）
	concurrentHttpRequestsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Subsystem: "sdk",
			Name:      "http_handle_concurrent",
			Help:      "Number of incoming HTTP Requests handling concurrently now",
		},
	)

	// 监控请求量，请求耗时等
	httpRequestHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: "sdk",
			Name:      "http_handle_requests",
			Help:      "Histogram statistics of http requests handle by elite http. Buckets by latency",
			Buckets:   []float64{0.001, 0.002, 0.005, 0.01, 0.05, 0.1, 0.2, 0.3, 0.4, 0.5, 0.8, 1, 2, 5, 10},
		},
		[]string{"code"},
	)

	summary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "test_summary",
			Help: "test summary",
			Objectives: map[float64]float64{
				0.5: 0.05,
			},
		}, []string{"name"},
	)

	completionTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_completion_timestamp_seconds",
		Help: "The timestamp of the last successful completion of a DB backup.",
	})

	// 用于测试cqh
	benchPress = prometheus.NewUntypedFunc(prometheus.UntypedOpts{
		Name: "bench_press",
	}, func() float64 {
		time.Sleep(time.Millisecond * 1)
		rand.Seed(time.Now().UnixNano())
		return rand.Float64()*40 + 100
	})

	deadLift = prometheus.NewUntypedFunc(prometheus.UntypedOpts{
		Name: "dead_lift",
	}, func() float64 {
		time.Sleep(time.Millisecond * 2)
		rand.Seed(time.Now().UnixNano())
		return rand.Float64()*60 + 100
	})

	deepSquat = prometheus.NewUntypedFunc(prometheus.UntypedOpts{
		Name: "deep_squat",
	}, func() float64 {
		time.Sleep(time.Millisecond * 3)
		rand.Seed(time.Now().UnixNano())
		return rand.Float64()*80 + 100
	})
)

func init() {
	//prometheus.MustRegister(httpRequestCount)
	//prometheus.MustRegister(concurrentHttpRequestsGauge)
	//prometheus.MustRegister(httpRequestHistogram)
	//prometheus.MustRegister(summary)
}

func main() {
	var count int64
	for {
		// 初始化一个pusher
		pusher := push.New("http://192.168.81.103:9091", "cqh")
		// 为pusher加入一些grouping key
		pusher.Grouping("instance", "test")

		// 向pusher中注册一个metric收集器
		// pusher.Collector(completionTime)

		// 向pusher中注册多个metrics
		registry := prometheus.NewRegistry()
		registry.MustRegister(benchPress, deadLift, deepSquat)

		// 将register添加进pusher
		pusher.Gatherer(registry)

		// 将各metrics中的指标推送到pushgateway
		pusher.Push()

		count++
		log.Printf("执行第 %d 次完成\n", count)
		time.Sleep(time.Second * 30)
	}
}

// GinMetricsMiddleware 在gin中间件中使用
func GinMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 统计接口请求数量
		httpRequestCount.Inc()

		// 监控并发量，进入接口前 +1
		concurrentHttpRequestsGauge.Inc()

		startTime := time.Now()

		// 后续处理逻辑
		c.Next()

		// after request
		finishTime := time.Now()

		// 监控计算接口耗时，请求数量等
		httpRequestHistogram.With(prometheus.Labels{"code": strconv.Itoa(http.StatusOK)}).Observe(float64(finishTime.Sub(startTime)) / (1000 * 1000 * 1000))

		// 监控并发量，离开接口后 -1
		concurrentHttpRequestsGauge.Dec()
	}
}

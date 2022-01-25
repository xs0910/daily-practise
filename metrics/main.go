package main

import (
	"expvar"
	"net/http"
	"runtime"
)

var (
	appleCounter      *expvar.Int
	GOMAXPROCSMetrics *expvar.Int
)

func init() {
	// 自定义新增两个字段
	appleCounter = expvar.NewInt("apple")
	GOMAXPROCSMetrics = expvar.NewInt("GOMAXPROCS")
	GOMAXPROCSMetrics.Set(int64(runtime.NumCPU()))

	// 继承接口 Var
	//upTimeMetric := &upTimeVar{value: time.Now().Local()}
	//expvar.Publish("uptime", upTimeMetric)
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		appleCounter.Add(1)
		_, _ = w.Write([]byte(`Go 语言编程之旅 `))
	})

	_ = http.ListenAndServe(":6060", http.DefaultServeMux)

	// 访问链接：http://127.0.0.1:6060/debug/vars
}

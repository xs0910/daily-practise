package gomock

// 安装GoMock
//$ go get github.com/golang/mock/gomock
//$ go install github.com/golang/mock/mockgen

// 在gomock目录下执行：
// mockgen -destination spider/mock/mock_spider.go -package spider daily-practise/gomock/spider Spider

// mockgen工具是 GoMock 提供的，用来 Mock 一个 Go 接口。它可以根据给定的接口，来自动生成 Mock 代码。有两种模式可以生成 Mock 代码，分别是源码模式和反射模式。
// 1. 源码模式
// mockgen -destination spider/mock/mock_spider.go -package spider -source spider/spider.go

// 2. 反射模式
// mockgen -destination spider/mock/mock_spider.go -package spider github.com/marmotedu/gopractise-demo/gomock/spider Spider

// 通过注释使用 mockgen
// 如果有多个文件，并且分散在不同的位置，那么我们要生成 Mock 文件的时候，需要对每个文件执行多次 mockgen 命令（这里假设包名不相同）。
// 这种操作还是比较繁琐的，mockgen 还提供了一种通过注释生成 Mock 文件的方式，此时需要借助go generate工具。
//go:generate mockgen -destination mock/mock_spider.go -package spider github.com/cz-it/blog/blog/Go/testing/gomock/example/spider Spider
// 在gomock目录下，执行以下命令，就可以自动生成 Mock 代码
// go generate ./...

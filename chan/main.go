package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	var urls []string
	for i := 0; i < 100; i++ {
		urls = append(urls, strconv.Itoa(i))
	}
	fmt.Println(time.Now())
	result := collect(urls)
	fmt.Println(time.Now())
	fmt.Println(result)
}

func collect(urls []string) []string {
	var result []string

	wg := &sync.WaitGroup{}
	response := make(chan string, 20)

	wgResponse := &sync.WaitGroup{}
	go func() {
		wgResponse.Add(1)

		for rc := range response {
			fmt.Println("response:", rc)
			result = append(result, rc)
		}
		wgResponse.Done()
	}()

	for _, url := range urls {
		wg.Add(1)
		go httpGet(url, response, wg)
	}

	// 等待结束
	wg.Wait()
	// 关闭response
	close(response)
	wgResponse.Wait()

	return result
}

func httpGet(url string, response chan string, wg *sync.WaitGroup) {
	// 也可以再传入一个channel，控制并发度
	defer wg.Done()

	time.Sleep(1 * time.Second)

	// 结果数据传入管道
	response <- fmt.Sprintf("http get:%s", url)
}

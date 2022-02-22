package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer1 := time.NewTimer(time.Second * 2)
	t1 := time.Now()
	fmt.Println("t1:", t1)

	t2 := <-timer1.C
	fmt.Println("t2:", t2)

	timer2 := time.NewTimer(time.Second * 2)
	<-timer2.C
	fmt.Println("2s后:", time.Now())

	time.Sleep(time.Second * 2)
	fmt.Println("再次2s后:", time.Now())

	<-time.After(time.Second * 2) // time.After 函数的返回值是chan time
	fmt.Println("再再次2s后:", time.Now())

	timer3 := time.NewTimer(time.Second)
	go func() {
		<-timer3.C
		fmt.Println("Timer3 expired:", time.Now())
	}()

	//stop := timer3.Stop() // 停止定时器
	//if stop {
	//	fmt.Println("Timer3 stopped")
	//}

	//fmt.Println("before:", time.Now())
	//timer4 := time.NewTimer(time.Second * 5)
	//timer4.Reset(time.Second * 1)
	//<-timer4.C
	//fmt.Println("after:", time.Now())
}

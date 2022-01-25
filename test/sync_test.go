package test

import (
	"bytes"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestSync(t *testing.T) {
	syncSample()
}

func syncSample() {
	p := sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	a := p.Get().(int)
	p.Put(1)
	b := p.Get().(int)
	fmt.Println(a, b)

	// Callers should not assume any relation between values passed to Put and the values returned by Get
	// 简单来说：就是get和put没有任何关系
	// Pool 就是为了减少GC压力的，重复利用内存，千万不要把他当做内存池使用
}

func TestSyncPool(t *testing.T) {
	syncPoolSample()
}

func syncPoolSample() {
	// 想在pool的基础上做一个限制池中对象数量的功能, 发现还是多次执行pool.New.
	// 期望是只执行一次NewBuffer，也就是只打印一次alloc。
	// 实际上每次执行，会打印多次alloc
	for i := 0; i < 10; i++ {
		// 多个协程想从pool中拿到对象
		go func() {
			c := getBuf()
			putBuf(c)
			log.Println("put done")
		}()
	}
	time.Sleep(3 * time.Second)
}

const MaxFrameSize = 2000

var bufPool = sync.Pool{
	New: func() interface{} {
		log.Println("alloc")
		return bytes.NewBuffer(make([]byte, 0, MaxFrameSize))
	},
}

var bufPoolChan = make(chan bool, 1)

func getBuf() *bytes.Buffer {
	bufPoolChan <- true
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	return b
}

func putBuf(b *bytes.Buffer) {
	bufPool.Put(b)
	<-bufPoolChan
}

// sync.Pool的源代码里说了，pool里的对象随时都有可能被自动移除，并且没有任何通知。sync.Pool的数量是不可控制的。

// Pool调用New与线程调度有关，Pool内部有一个localPool的数组，每个P对应其中一个localPool，
// 在当前P执行goroutine的时候，优先从当前的localPool的private变量取，取不到在从shared列表里面取，再取不到就尝试从别的P的localPool的shared里面偷一个。
// 最后实在取不到就New一个。

// 由于你的bufPoolChan限制基本上10个goroutine就在两个P后面排队轮流执行，所以alloc就会出现两次，后面的基本就是从这两个localPool的private取出来的。
// 如果取消这个限制，10个goroutine很快就被分配到10个P上去了，对应就有10个localPool，10次每次取private都取不到，
// 取shared列表也取不到，别的localPool也没得偷，就会New10次，alloc就会出现10次。

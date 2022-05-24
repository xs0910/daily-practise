package kafka_sarama_demo

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strconv"
	"testing"
	"time"
)

func TestProducer(t *testing.T) {
	fmt.Printf("producer_test\n")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel 返回
	config.Producer.Return.Errors = true
	config.Version = sarama.V2_1_0_0

	// 异步连接kafka
	producer, err := sarama.NewAsyncProducer([]string{"192.168.81.110:9092"}, config)
	if err != nil {
		fmt.Printf("producer_test create producer error :%s\n", err.Error())
		return
	}

	defer producer.AsyncClose()

	// send message 构造一个消息
	msg := &sarama.ProducerMessage{
		Topic: "kafka_go_test",
		Key:   sarama.StringEncoder("go_test"),
	}

	value := "this is message: "
	n := 0
	for {
		n++
		message := value + strconv.Itoa(n)
		msg.Value = sarama.ByteEncoder(message)
		fmt.Printf("input [%s]\n", message)

		// send to chain
		producer.Input() <- msg

		select {
		case suc := <-producer.Successes():
			fmt.Printf("offset: %d,  timestamp: %s", suc.Offset, suc.Timestamp.String())
		case fail := <-producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		}

		time.Sleep(time.Second * 5)
	}
}

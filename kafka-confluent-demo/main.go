package main

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const Topic = "testTopic"
const Broker = "192.168.81.110:9092"
const NumPartition = 2
const ReplicationFactor = 1
const ConsumerGroup11 = "consumerTest1"

func main() {
	fmt.Println("Kafka demo start")
	// 创建topic
	KafkaCreateTopic()
	// 创建生产者
	go KafkaProducer()
	// 创建消费者
	go KafkaConsumer("group1")
	go KafkaConsumer("group2")
	go KafkaConsumer("group3")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating \n", sig)
			run = false
		}
	}
}

// KafkaCreateTopic 创建topic
func KafkaCreateTopic() {
	client, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": Broker,
	})
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}
	defer client.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 在集群上创建主题
	// 设置管理员选项以等待操作完成
	maxDur, err := time.ParseDuration("60s")
	if err != nil {
		panic("ParseDuration(60s)")
	}

	results, err := client.CreateTopics(
		ctx,
		// 通过TopicSpecification结构可以同时创建多个主题
		[]kafka.TopicSpecification{
			{
				Topic:             Topic,
				NumPartitions:     NumPartition,
				ReplicationFactor: ReplicationFactor,
			},
		},
		kafka.SetAdminRequestTimeout(maxDur),
	)
	if err != nil {
		fmt.Printf("Failed to create topic: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Print results")
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}
}

// KafkaProducer 消息生产者
func KafkaProducer() {
	topic := Topic
	broker := Broker
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Producer %v\n", producer)

	// Optional delivery channel, if not specified the Producer object's Events channel is used.
	deliveryChan := make(chan kafka.Event)

	// 每5秒向kafka发送一条消息
	n := 0
	for {
		n++
		value := strconv.Itoa(n) + " Hello Go!"
		err = producer.Produce(
			&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &topic,
					Partition: kafka.PartitionAny,
				},
				Value: []byte(value),
				Headers: []kafka.Header{
					{
						Key:   "myTestHeader",
						Value: []byte("Header values are binary"),
					},
				},
			},
			deliveryChan,
		)

		e := <-deliveryChan
		m := e.(*kafka.Message)

		if m.TopicPartition.Error != nil {
			fmt.Printf("生产者：Delivery Failed： %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("生产者：Deliveried message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}

		time.Sleep(5 * time.Second)
	}

	close(deliveryChan)
}

func KafkaConsumer(consumerGroup string) {
	broker := Broker
	group := consumerGroup
	topics := Topic
	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     broker,
		"broker.address.family": "v4",
		"group.id":              group,
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", consumer)

	err = consumer.SubscribeTopics([]string{topics}, nil)
	run := true

	for run {
		select {
		case sig := <-signChan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("消费者consumerGroup: %s Message on %s:\n%s\n",
					group, e.TopicPartition, string(e.Value))
				if e.Headers != nil {
					fmt.Printf("Headers: %v\n", e.Headers)
				}
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	consumer.Close()
}

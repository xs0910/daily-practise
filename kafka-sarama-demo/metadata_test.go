package kafka_sarama_demo

import (
	"fmt"
	"github.com/Shopify/sarama"
	"testing"
)

func TestMetadata(t *testing.T) {
	fmt.Printf("metadata test\n")

	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0

	client, err := sarama.NewClient([]string{"192.168.81.110:9092"}, config)
	if err != nil {
		fmt.Printf("metadata_test try create client err: %s\n", err.Error())
		return
	}

	defer client.Closed()

	// get topic set
	topics, err := client.Topics()
	if err != nil {
		fmt.Printf("try get topics err %s\n", err.Error())
		return
	}

	fmt.Printf("topics(%d):\n", len(topics))
	for _, topic := range topics {
		fmt.Println("    ", topic)
	}

	// get broker set
	brokers := client.Brokers()
	fmt.Printf("broker set(%d):\n", len(brokers))
	for _, broker := range brokers {
		fmt.Printf("    %s\n", broker.Addr())
	}
}

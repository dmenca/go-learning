package main

import (
	"fmt"
	"log"
	"time"

	"encoding/json"

	"github.com/Shopify/sarama"
)

// MyMessage 代表Kafka消息的数据结构
type MyMessage struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

func main() {
	// kafka broker地址
	brokers := []string{"10.111.32.63:9092"}

	// 创建一个kafka的配置
	config := sarama.NewConfig()
	// 配置sasl
	config.Producer.Return.Successes = true
	config.Net.SASL.Enable = true
	config.Net.SASL.User = "admin"
	config.Net.SASL.Password = "wrW4sYXJEvAJ46s"
	config.Net.SASL.Mechanism = sarama.SASLTypePlaintext

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalln("Error create producer:", err)
	}

	topic := "test-ca"
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln("Error closing consumer:", err)
		}
	}()

	for i := 0; i < 10; i++ {
		myMessage := &MyMessage{ID: i, Content: fmt.Sprintf("Message %d", i)}
		myMessageBytes, err := json.Marshal(myMessage)

		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(string(myMessageBytes)),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Printf("Error producing message: %s\n", err)
		} else {
			fmt.Printf("Produced message to partition %d at offset %d\n", partition, offset)
		}
		time.Sleep(1 * time.Second)
	}

}

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

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

	config.Net.SASL.Enable = true
	config.Net.SASL.User = "admin"
	config.Net.SASL.Password = "wrW4sYXJEvAJ46s"
	config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalln("Error create consumer:", err)
	}

	topic := "test-ca"
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln("Error closing consumer:", err)
		}
	}()

	// 获取partitions列表 返回[]int32数组 分区的id列表
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatalln("Error get partitions:", err)
	}

	for _, partition := range partitions {
		// 根据parttion topic进行消费
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			log.Fatalln("Error starting partition consumer:", err)
		}
		defer func() {
			if err := pc.Close(); err != nil {
				log.Fatalln("Error closing consumer:", err)
			}
		}()

		// 创建一个线程消费一个partition
		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				var myMessage MyMessage
				err := json.Unmarshal(message.Value, &myMessage)
				if err != nil {
					log.Printf("Error unmarshaling message: %v", err)
				} else {
					fmt.Printf("Received message: %+v\n", myMessage)
				}
				fmt.Printf("Partition: %d, Offset: %d, Key: %s, Value: %s\n",
					message.Partition, message.Offset, string(message.Key), string(message.Value))
			}
		}(pc)
	}

	// 等待结束
	select {}

}

package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	// kafka broker地址
	brokers := []string{"10.111.32.63:9092"}

	// 创建一个kafka的配置
	config := sarama.NewConfig()
	// 配置sasl
	saslConfig := config.Net.SASL
	saslConfig.Enable = true
	saslConfig.User = "admin"
	saslConfig.Password = "wrW4sYXJEvAJ46s"
	saslConfig.Mechanism = "SASL_PLAINTEXT"
	fmt.Println(saslConfig)
	fmt.Println(brokers)
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalln("Error create consumer:", err)
	}

	topic := "WhaleCollect"

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln("Error closing consumer:", err)
		}
	}()

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatalln("Error get partitions:", err)
	}

	for _, partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatalln("Error starting partition consumer:", err)
		}
		defer func() {
			if err := pc.Close(); err != nil {
				log.Fatalln("Error closing consumer:", err)
			}
		}()
		// 创建一个线程消费

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				fmt.Printf("Partition: %d, Offset: %d, Key: %s, Value: %s\n",
					message.Partition, message.Offset, string(message.Key), string(message.Value))
			}
		}(pc)
	}

	// 等待结束
	select {}

}

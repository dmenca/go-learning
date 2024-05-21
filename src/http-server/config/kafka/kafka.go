package kafka

import "github.com/Shopify/sarama"

type Config struct {
	Brokers  []string
	KafkaCfg *sarama.Config
}

func (c *Config) GetClient() (sarama.Client, error) {
	client, err := sarama.NewClient(c.Brokers, c.KafkaCfg)
	if err != nil {
		return nil, err
	}
	return client, err
}

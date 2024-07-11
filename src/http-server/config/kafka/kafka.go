package kafka

import (
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Brokers  []string
	KafkaCfg *sarama.Config
}

func NewDefaultConfig() *Config {
	kafkaCfg := sarama.NewConfig()
	kafkaCfg.Producer.Return.Successes = true
	return &Config{
		Brokers:  []string{"localhost:9092"},
		KafkaCfg: kafkaCfg,
	}
}

func (c *Config) GetClient() (sarama.Client, error) {
	client, err := sarama.NewClient(c.Brokers, c.KafkaCfg)
	if err != nil {
		return nil, err
	}
	return client, err
}

func (c *Config) ProduceMsgByID(msgBytes []byte, Id string, topic string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(Id),
		Value: sarama.StringEncoder(msgBytes),
	}

	// 连接kafka
	producer, err := sarama.NewSyncProducer(c.Brokers, c.KafkaCfg)
	if err != nil {
		log.Errorf("ProduceMsgByID() failed to produce sync msg, error: %v", err)
		return err
	}
	defer func(client sarama.SyncProducer) error {
		err := client.Close()
		if err != nil {
			log.Errorf("ProduceMsgByID() producer closed with err:%v", err)
			return err
		}
		log.Info("ProduceMsgByID() producer closed succeed")
		return nil
	}(producer)
	// 发送消息

	log.Debugf("ProduceMsgByID() kafka config: %v", c)

	pid, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Errorf("ProduceMsgByID() producer send msg failed, err:%v", err)
		return err
	}
	log.Debugf("ProduceMsgByID() producer send msg succeed ,pid:%v offset:%v\n", pid, offset)
	return nil
}

func (c *Config) ProduceMsgs(msgs []*sarama.ProducerMessage) error {
	log.Debugf("ProduceMsgs() kafka config: %+#v", c)
	// 连接kafka
	producer, err := sarama.NewSyncProducer(c.Brokers, c.KafkaCfg)
	if err != nil {
		log.Errorf("ProduceMsgs() failed to produce sync msg, error: %v", err)
		return err
	}
	defer func(client sarama.SyncProducer) error {
		err := client.Close()
		if err != nil {
			log.Errorf("ProduceMsgs() producer closed with err:%v", err)
			return err
		}
		log.Info("ProduceMsgs() producer closed succeed")
		return nil
	}(producer)

	// 发送消息
	err1 := producer.SendMessages(msgs)
	if err1 != nil {
		log.Errorf("ProduceMsgs() producer send msg failed, err:%v", err1)
		return err1
	}
	log.Debugf("ProduceMsgs() producer send all the msgs succeed !")
	return nil
}

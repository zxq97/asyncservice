package kafka

import (
	"asyncservice/conf"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

var (
	client *sarama.SyncProducer
)

const (
	DefaultDialTimeout  = 500 * time.Millisecond
	DefaultReadTimeout  = 5 * time.Second
	DefaultWriteTimeout = 5 * time.Second
)

func InitClient(kafkaConf conf.KafkaConf) error {
	kfkConf := sarama.NewConfig()
	kfkConf.Producer.RequiredAcks = sarama.WaitForAll
	kfkConf.Producer.Retry.Max = 3
	kfkConf.Producer.Return.Successes = true
	kfkConf.Net.DialTimeout = DefaultDialTimeout
	kfkConf.Net.ReadTimeout = DefaultReadTimeout
	kfkConf.Net.WriteTimeout = DefaultWriteTimeout
	producer, err := sarama.NewSyncProducer([]string{kafkaConf.Addr}, kfkConf)
	if err != nil {
		log.Printf("get producer err: %v", err)
		return err
	}
	client = &producer
	return err
}

func SendMessage(topic string, key []byte, data []byte) error {
	partition, offset, err := (*client).SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(data),
	})
	if err != nil {
		log.Printf("kfk.sendmessage: key: %s data: %s, partition: %d, offset: %d, err: %s", key, data, partition, offset, err)
		return err
	}
	log.Printf("SendMessage: info, topic=%v key=%v data=%v p=%v offset=%v", topic, string(key), string(data), partition, offset)
	return nil
}

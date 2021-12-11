package kafka

import (
	"github.com/Shopify/sarama"
	"log"
)

func NewKafkaConsumer(broker string) (sarama.Consumer, error) {
	consumer, err := sarama.NewConsumer([]string{broker}, nil)
	if err != nil {
		log.Println("consumer connect err", err)
		return nil, err
	}
	return consumer, nil
}

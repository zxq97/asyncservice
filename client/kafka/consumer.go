package kafka

import (
	"asyncservice/global"
	"github.com/Shopify/sarama"
)

func NewKafkaConsumer(broker []string) (sarama.Consumer, error) {
	consumer, err := sarama.NewConsumer(broker, nil)
	if err != nil {
		global.ExcLog.Println("consumer connect err", err)
		return nil, err
	}
	return consumer, nil
}

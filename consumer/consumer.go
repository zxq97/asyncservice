package consumer

import (
	"asyncservice/client/kafka"
	"asyncservice/consumer/article"
	"asyncservice/consumer/social"
	"asyncservice/global"
	"encoding/json"
	"github.com/Shopify/sarama"
)

func InitConsumer(broker []string, topic string) {
	consumer, err := kafka.NewKafkaConsumer(broker)
	defer consumer.Close()
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		return
	}

	for _, p := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(topic, p, sarama.OffsetNewest)
		if err != nil {
			global.ExcLog.Printf("partitionconsumer err %v", err)
			continue
		}

		for m := range partitionConsumer.Messages() {
			process(m)
		}
	}
}

func process(message *sarama.ConsumerMessage) {
	val := message.Value
	global.ExcLog.Printf("ProcessComics: info, key %v value %v", string(message.Key), string(message.Value))
	event := new(kafka.KafkaMessage)
	err := json.Unmarshal(val, event)
	if err != nil {
		global.ExcLog.Printf("process json unmarshal %v err %v", string(message.Value), err)
		return
	}
	switch event.Event {
	case kafka.EventPublish:
		article.PublishArticle(event)
	case kafka.EventFollow:
		social.Follow(event)
	case kafka.EventUnfollow:
		social.Unfollow(event)
	}

}

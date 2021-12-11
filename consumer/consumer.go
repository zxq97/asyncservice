package consumer

import (
	"asyncservice/client/kafka"
	"asyncservice/consumer/article"
	"asyncservice/consumer/social"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

func InitConsumer(broker, topic string) {
	consumer, err := kafka.NewKafkaConsumer(broker)
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		return
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg, ok := <-partitionConsumer.Messages():
			if ok {
				process(msg)
			}
		case err, ok := <-partitionConsumer.Errors():
			if ok {
				fmt.Println(err)
			}
		}
	}
}

func process(message *sarama.ConsumerMessage) {
	val := message.Value
	log.Printf("ProcessComics: info, key %v value %v", string(message.Key), string(message.Value))
	event := new(kafka.KafkaMessage)
	err := json.Unmarshal(val, event)
	if err != nil {
		log.Printf("process json unmarshal %v err %v", string(message.Value), err)
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

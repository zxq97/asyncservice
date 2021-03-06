package consumer

import (
	"asyncservice/client/kafka"
	"asyncservice/consumer/article"
	"asyncservice/consumer/social"
	"asyncservice/global"
	"asyncservice/util/concurrent"
	"context"
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

	wg := concurrent.NewWaitGroup()

	for _, p := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(topic, p, sarama.OffsetNewest)
		if err != nil {
			global.ExcLog.Printf("partitionconsumer err %v", err)
			continue
		}

		wg.Run(func() {
			for m := range partitionConsumer.Messages() {
				process(m)
			}
		})
	}
	wg.Wait()
}

func process(message *sarama.ConsumerMessage) {
	val := message.Value
	global.InfoLog.Printf("ProcessComics: info, key %v value %v", string(message.Key), string(message.Value))
	event := new(kafka.KafkaMessage)
	err := json.Unmarshal(val, event)
	if err != nil {
		global.ExcLog.Printf("process json unmarshal %v err %v", string(message.Value), err)
		return
	}
	// fixme context需要换
	ctx := context.TODO()
	switch event.Event {
	case kafka.EventPublish:
		article.PublishArticle(ctx, event)
	case kafka.EventFollow:
		social.Follow(ctx, event)
	case kafka.EventUnfollow:
		social.Unfollow(ctx, event)
	}

}

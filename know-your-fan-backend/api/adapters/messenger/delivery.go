package messenger

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (k *KafkaProducer) handleDeliveryReport() {
	for e := range k.deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				log.Printf("Delivery failed: %v\n", ev.TopicPartition.Error)
			} else {
				log.Printf("Message delivered to %v\n", ev.TopicPartition)
			}
		}
	}
}

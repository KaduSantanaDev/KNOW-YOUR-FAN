package messenger

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaProducer struct {
	producer     *kafka.Producer
	deliveryChan chan kafka.Event
}

func NewKafkaProducer(brokers string) *KafkaProducer {
	config := &kafka.ConfigMap{
		"bootstrap.servers":   brokers,
		"delivery.timeout.ms": "0",
		"acks":                "all",
		"enable.idempotence":  "true",
	}

	p, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}

	k := &KafkaProducer{
		producer:     p,
		deliveryChan: make(chan kafka.Event),
	}

	go k.handleDeliveryReport()
	return k
}

func (k *KafkaProducer) Publish(message, topic string, key []byte) error {
	msg := &kafka.Message{
		Value:          []byte(message),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	return k.producer.Produce(msg, k.deliveryChan)
}

func (k *KafkaProducer) Close() {
	k.producer.Flush(3000)
	close(k.deliveryChan)
}

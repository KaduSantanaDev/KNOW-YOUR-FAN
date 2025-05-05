package messenger

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewKafkaConsumer(brokers, topic string) *KafkaConsumer {
	config := &kafka.ConfigMap{
		"bootstrap.servers":  brokers,
		"group.id":           "document-validator-group",
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	}

	c, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatalf("fail to create a consumer: %v", err)
	}

	if err := c.Subscribe(topic, nil); err != nil {
		log.Fatalf("fail to subcribge topics: %v", err)
	}

	return &KafkaConsumer{consumer: c}
}

func (kc *KafkaConsumer) Consume(ctx context.Context, handle func(msg *kafka.Message)) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	log.Println("start consuming")

	for {
		select {
		case <-ctx.Done():
			log.Println("Context canceled")
			return
		case sig := <-sigchan:
			log.Printf("Sinal capturado: %v, stop consuming", sig)
			return
		default:
			msg, err := kc.consumer.ReadMessage(-1)
			if err != nil {
				log.Printf("errorto read messages, %v", err)
				continue
			}

			log.Printf("📨 Message recieved (key=%s): %s", string(msg.Key), string(msg.Value))
			handle(msg)

			_, err = kc.consumer.CommitMessage(msg)
			if err != nil {
				log.Printf("error to commit offset: %v", err)
			}
		}
	}
}

func (kc *KafkaConsumer) Close() {
	kc.consumer.Close()
	log.Println("Consumer finished")
}

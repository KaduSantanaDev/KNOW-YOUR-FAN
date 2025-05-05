package service

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"context"

	"github.com/KaduSantanaDev/know-your-fan-api/adapters/database"
	"github.com/KaduSantanaDev/know-your-fan-api/adapters/messenger"
	"github.com/KaduSantanaDev/know-your-fan-api/application"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type ClientService struct {
	Repository database.ClientDB
}

func NewClientService(repository database.ClientDB) *ClientService {
	return &ClientService{
		Repository: repository,
	}
}

func (c *ClientService) Create(client application.ClientInterface) (application.ClientInterface, error) {
	newClient := application.NewClient()
	newClient.Address = client.GetAddress()
	newClient.Email = client.GetEmail()
	newClient.Name = client.GetName()
	newClient.CPF = client.GetCPF()
	newClient.Document = client.GetDocument()

	result, err := c.Repository.Create(newClient)
	if err != nil {
		return nil, err
	}

	if err := c.sendMessage(newClient.GetID(), newClient.GetName(), newClient.GetDocument()); err != nil {
		return nil, err
	}

	if !c.consumeMessage(newClient.GetID(), newClient.GetName()) {
		return nil, errors.New("Invalid document")
	}

	return result, nil
}

func (c *ClientService) sendMessage(id, name string, document []byte) error {
	event := messenger.ClientCreatedEvent{
		ID:       id,
		Name:     name,
		Document: document,
		Valid:    nil,
	}

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	producer := messenger.NewKafkaProducer("kafka:9092")
	defer producer.Close()

	if err := producer.Publish(string(data), "document-validation", []byte("RG")); err != nil {
		log.Println("ERRO: ")
		log.Fatalf(err.Error())
		return err
	}

	return nil
}

func (c *ClientService) consumeMessage(id, name string) bool {
	consumer := messenger.NewKafkaConsumer("kafka:9092", "document-validation-result")
	resultChan := make(chan bool)

	go func() {
		defer close(resultChan)

		consumer.Consume(context.Background(), func(msg *kafka.Message) {
			var result messenger.RecieveClientEvent
			if err := json.Unmarshal(msg.Value, &result); err != nil {
				log.Printf(err.Error())
				return
			}
			if string(msg.Key) != id {
				return
			}

			if result.Valid != nil {
				resultChan <- *result.Valid
			}
		})
	}()

	select {
	case valid := <-resultChan:
		return valid
	case <-time.After(5 * time.Second):
		return false
	}

}

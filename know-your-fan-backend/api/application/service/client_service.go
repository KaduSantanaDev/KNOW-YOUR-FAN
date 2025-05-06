package service

import (
	"encoding/json"
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

func (c *ClientService) GetAll() ([]application.ClientInterface, error) {
	return c.Repository.GetAll()
}

func (c *ClientService) Create(client application.ClientInterface) (application.ClientInterface, error) {
	newClient := application.Client{
		ID:       client.GetID(),
		Address:  client.GetAddress(),
		Email:    client.GetEmail(),
		Name:     client.GetName(),
		CPF:      client.GetCPF(),
		Document: client.GetDocument(),
	}

	createdClient, err := c.Repository.Create(&newClient)
	if err != nil {
		return nil, err
	}

	if err := c.sendMessage(newClient.GetID(), newClient.GetName(), newClient.GetDocument()); err != nil {
		return nil, err
	}

	c.consumeMessage(newClient.GetID(), newClient.GetName())

	return createdClient, nil
}

func (c *ClientService) UpdateStatus(client application.ClientInterface) (application.ClientInterface, error) {
	client.UpdateStatus(true)

	result, err := c.Repository.UpdateStatus(client)
	if err != nil {
		return nil, err
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
	client, err := c.Repository.GetByID(id)
	if err != nil {
		return false
	}
	log.Println("Achou pelo id")

	go func() {
		defer close(resultChan)

		consumer.Consume(context.Background(), func(msg *kafka.Message) {
			var result messenger.RecieveClientEvent
			if err := json.Unmarshal(msg.Value, &result); err != nil {
				log.Printf(err.Error())
				return
			}
			log.Println("Consumiu a mensagem", result.Valid)

			if result.Valid != nil {
				client.UpdateStatus(true)
				log.Println("Client updated")

				_, err := c.Repository.UpdateStatus(client)
				if err != nil {
					return
				}
				resultChan <- *result.Valid
				log.Println("Foi pro channel")
			}
		})
	}()

	select {
	case valid := <-resultChan:
		return valid
	case <-time.After(15 * time.Second):
		return false
	}

}

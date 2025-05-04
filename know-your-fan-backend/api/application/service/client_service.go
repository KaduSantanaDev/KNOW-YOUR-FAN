package service

import (
	"encoding/json"
	"log"

	"github.com/KaduSantanaDev/know-your-fan-api/adapters/database"
	"github.com/KaduSantanaDev/know-your-fan-api/adapters/messenger"
	"github.com/KaduSantanaDev/know-your-fan-api/application"
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
	return result, nil
}

func (c *ClientService) SendMessage(id, name string, document []byte) error {
	event := messenger.ClientCreatedEvent{
		ID:       id,
		Name:     name,
		Document: document,
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

	log.Println("mandou a mensagem")
	return nil
}

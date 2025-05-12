package service

import (
	"encoding/json"
	"log"

	"github.com/KaduSantanaDev/document-validation-api/adapters/database"
	"github.com/KaduSantanaDev/document-validation-api/adapters/messenger"
	"github.com/KaduSantanaDev/document-validation-api/application"
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

	if err := c.sendMessage(createdClient.GetID(), createdClient.GetName(), createdClient.GetDocument()); err != nil {
		return nil, err
	}

	return createdClient, nil
}

func (c *ClientService) sendMessage(id, name string, document []byte) error {
	event := messenger.ClientCreatedEvent{
		ID:       id,
		Name:     name,
		Document: document,
		Valid:    false,
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

func (c *ClientService) GetByID(id string) (application.ClientInterface, error) {
	return c.Repository.GetByID(id)
}

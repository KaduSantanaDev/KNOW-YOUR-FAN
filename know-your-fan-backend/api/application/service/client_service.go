package service

import (
	"github.com/KaduSantanaDev/know-your-fan-api/adapters/database"
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

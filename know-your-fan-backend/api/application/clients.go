package application

import (
	"github.com/google/uuid"
)

type Client struct {
	ID       string
	Name     string
	Email    string
	CPF      string
	Document []byte
	Address  Address
}

type ClientInterface interface {
	GetID() string
	GetEmail() string
	GetDocument() string
	GetAddress() Address
}

type ClientServiceInterface interface {
	Create(name, email, cpf string, address Address)
}

func (c *Client) GetID() string {
	return c.ID
}
func (c *Client) GetName() string {
	return c.Name
}
func (c *Client) GetEmail() string {
	return c.Email
}
func (c *Client) GetCPF() string {
	return c.CPF
}
func (c *Client) GetAddress() Address {
	return c.Address
}

func NewClient() *Client {
	return &Client{
		ID: uuid.New().String(),
	}
}

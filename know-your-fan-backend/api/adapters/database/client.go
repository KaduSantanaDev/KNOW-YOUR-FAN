package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/KaduSantanaDev/know-your-fan-api/application"
)

type ClientDB struct {
	db *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB {
	return &ClientDB{db: db}
}

func (c *ClientDB) GetAll() ([]application.ClientInterface, error) {
	rows, err := c.db.Query(`SELECT id, name, email, cpf, document, street, number, complement, neighborhood, city, state, cep, status FROM clients`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []application.ClientInterface

	for rows.Next() {
		var client application.Client
		var address application.Address
		err := rows.Scan(
			&client.ID,
			&client.Name,
			&client.Email,
			&client.CPF,
			&client.Document,
			&address.Street,
			&address.Number,
			&address.Complement,
			&address.Neighborhood,
			&address.City,
			&address.State,
			&address.CEP,
			&client.Status,
		)
		if err != nil {
			return nil, err
		}
		client.Address = address
		clients = append(clients, &client)
	}

	return clients, nil
}

func (c *ClientDB) GetByID(id string) (application.ClientInterface, error) {
	row := c.db.QueryRow(`SELECT id, name, email, cpf, document, street, number, complement, neighborhood, city, state, cep, status FROM clients WHERE id = $1`, id)

	var client application.Client
	var address application.Address

	err := row.Scan(
		&client.ID,
		&client.Name,
		&client.Email,
		&client.CPF,
		&client.Document,
		&address.Street,
		&address.Number,
		&address.Complement,
		&address.Neighborhood,
		&address.City,
		&address.State,
		&address.CEP,
		&client.Status,
	)
	if err != nil {
		return nil, err
	}

	client.Address = address

	return &client, nil
}

func (c *ClientDB) Create(client application.ClientInterface) (application.ClientInterface, error) {
	if err := c.validateClientDoesNotExist(client); err != nil {
		return nil, err
	}

	stmt, err := c.db.Prepare(`
		INSERT INTO clients(id, name, email, cpf, document, street, number, complement, neighborhood, city, state, cep, status)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		client.GetID(),
		client.GetName(),
		client.GetEmail(),
		client.GetCPF(),
		client.GetDocument(),
		client.GetAddress().Street,
		client.GetAddress().Number,
		client.GetAddress().Complement,
		client.GetAddress().Neighborhood,
		client.GetAddress().City,
		client.GetAddress().State,
		client.GetAddress().CEP,
		client.GetStatus(),
	)

	if err != nil {
		log.Println("Erro ao inserir cliente:", err)
		return nil, err
	}

	return client, nil
}

func (c *ClientDB) validateClientDoesNotExist(client application.ClientInterface) error {
	var existingID string
	err := c.db.QueryRow(`SELECT id FROM clients WHERE id = $1`, client.GetID()).Scan(&existingID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return fmt.Errorf("erro ao verificar cliente no banco: %w", err)
	}

	return fmt.Errorf("cliente com ID %s já existe", client.GetID())
}

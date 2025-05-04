package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/KaduSantanaDev/know-your-fan-api/application"
)

type ClientDB struct {
	db *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB {
	return &ClientDB{db: db}
}

func (c *ClientDB) Create(client application.ClientInterface) (application.ClientInterface, error) {
	if err := c.validateClientDoesNotExist(client); err != nil {
		return nil, err
	}

	stmt, err := c.db.Prepare(`insert into clients(id, name, email, cpf, document, street, number, complement, neighborhood, city, state, cep)
									values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
		`)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		client.GetID(),
		client.GetName(),
		client.GetCPF(),
		client.GetEmail(),
		client.GetDocument(),
		client.GetAddress().Street,
		client.GetAddress().Number,
		client.GetAddress().Complement,
		client.GetAddress().Neighborhood,
		client.GetAddress().City,
		client.GetAddress().State,
		client.GetAddress().CEP,
	)

	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *ClientDB) validateClientDoesNotExist(client application.ClientInterface) error {
	var existingID string
	err := c.db.QueryRow(`SELECT id FROM clients WHERE id = ?`, client.GetID()).Scan(&existingID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return fmt.Errorf("erro ao verificar cliente no banco: %w", err)
	}

	return fmt.Errorf("cliente com ID %s já existe", client.GetID())
}

package handlers

type RetrieveClientDTO struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	CPF    string `json:"cpf"`
	Status bool   `json:"status"`
}

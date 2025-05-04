package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KaduSantanaDev/know-your-fan-api/application"
	"github.com/KaduSantanaDev/know-your-fan-api/application/service"
)

type ClientHandler struct {
	ClientService service.ClientService
}

func NewClientHandler(clientService service.ClientService) *ClientHandler {
	return &ClientHandler{
		ClientService: clientService,
	}
}

func (c *ClientHandler) Create(w http.ResponseWriter, r *http.Request) {
	var createClientDTO CreateClientDTO

	if err := json.NewDecoder(r.Body).Decode(&createClientDTO); err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newClient := application.NewClient()
	newClient.Name = createClientDTO.Name
	newClient.Email = createClientDTO.Email
	newClient.CPF = createClientDTO.CPF
	newClient.Document = createClientDTO.Document
	newClient.Address = application.Address{
		Street:       createClientDTO.Street,
		Number:       createClientDTO.Number,
		Complement:   createClientDTO.Complement,
		Neighborhood: createClientDTO.Neighborhood,
		City:         createClientDTO.City,
		State:        createClientDTO.State,
		CEP:          createClientDTO.CEP,
	}

	c.ClientService.Create(newClient)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "client created",
	})
}

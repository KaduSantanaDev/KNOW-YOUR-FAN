package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/KaduSantanaDev/document-validation-api/application"
	"github.com/KaduSantanaDev/document-validation-api/application/service"
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
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	documentFile, _, err := r.FormFile("document")
	if err != nil {
		http.Error(w, "Documento inválido", http.StatusBadRequest)
		return
	}

	defer documentFile.Close()

	docBytes, _ := io.ReadAll(documentFile)

	newClient := application.NewClient()
	newClient.Name = r.FormValue("name")
	newClient.Email = r.FormValue("email")
	newClient.CPF = r.FormValue("cpf")
	newClient.Document = docBytes
	newClient.Address = application.Address{
		Street:       r.FormValue("street"),
		Number:       parseInt32(r.FormValue("number")),
		Complement:   r.FormValue("complement"),
		Neighborhood: r.FormValue("neighborhood"),
		City:         r.FormValue("city"),
		State:        r.FormValue("state"),
		CEP:          r.FormValue("cep"),
	}
	newClient.Status = false

	createdClient, err := c.ClientService.Create(newClient)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "error creating client",
			"error":   err.Error(),
		})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "error getting client",
			"error":   err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "client created",
		"client":  createdClient,
	})
}

func parseInt32(s string) int32 {
	i, _ := strconv.Atoi(s)
	return int32(i)
}

func (c *ClientHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	clients, err := c.ClientService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "error getting clients",
			"error":   err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "clients retrieved",
		"clients": clients,
	})
}

func (c *ClientHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	client, err := c.ClientService.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "error getting client",
			"error":   err.Error(),
		})
		return
	}

	clientDTO := RetrieveClientDTO{
		Name:   client.GetName(),
		Email:  client.GetEmail(),
		CPF:    client.GetCPF(),
		Status: client.GetStatus(),
	}

	if client.GetStatus() {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "client retrieved",
			"client":  clientDTO,
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]any{
		"message": "client must be activated",
		"client":  clientDTO,
	})
}

package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/KaduSantanaDev/know-your-fan-api/adapters/database"
	handlers "github.com/KaduSantanaDev/know-your-fan-api/adapters/http"
	"github.com/KaduSantanaDev/know-your-fan-api/application/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://root:secret@postgres:5432/clients?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	clientDB := database.NewClientDB(db)
	clientService := service.NewClientService(*clientDB)
	clientHandler := handlers.NewClientHandler(*clientService)

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	}))
	r.Use(middleware.Logger)

	r.Post("/api/v1/clients", clientHandler.Create)
	r.Get("/api/v1/clients", clientHandler.GetAll)

	log.Println("Servidor ouvindo em http://localhost:3031")
	if err := http.ListenAndServe(":3031", r); err != nil {
		log.Fatal(err)
	}
}

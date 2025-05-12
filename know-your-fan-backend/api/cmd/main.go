package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KaduSantanaDev/document-validation-api/adapters/database"
	handlers "github.com/KaduSantanaDev/document-validation-api/adapters/http"
	"github.com/KaduSantanaDev/document-validation-api/application/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis do sistema", err)
	}
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dsn)
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
	r.Get("/api/v1/clients", clientHandler.GetByID)

	log.Println("Servidor ouvindo em http://localhost:3031")
	if err := http.ListenAndServe(":3031", r); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/igosantana/igo-refrigeration-api/internal/config"
	"github.com/igosantana/igo-refrigeration-api/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	config.LoadEnv()

	dbConn := db.Connect()

	r := routes.SetupRoutes(dbConn)

	log.Println("API rodando na porta 8080")
	http.ListenAndServe(":8080", r)
}

package db

import (
	"context"
	"log"
	"time"

	"github.com/igosantana/igo-refrigeration-api/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() *pgxpool.Pool {
	dsn := config.AppConfig.DatabaseURL

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Erro ao criar pool de conexão: %v", err)
	}

	if err := dbpool.Ping(ctx); err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
	}

	log.Println("Conectado ao banco com sucesso")

	return dbpool
}

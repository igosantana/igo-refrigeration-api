package db

import (
	"log"

	"github.com/igosantana/igo-refrigeration-api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := config.AppConfig.DatabaseURL

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados com GORM: %v", err)
	}

	log.Println("✅ Conectado ao banco com GORM com sucesso")

	return db
}

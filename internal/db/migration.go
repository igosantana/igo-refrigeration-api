package db

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/igosantana/igo-refrigeration-api/internal/db/migrations"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	allMigrations := migrations.GetAllMigrations()

	m := gormigrate.New(db, gormigrate.DefaultOptions, allMigrations)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Erro ao rodar as migrations: %v", err)
	}
}

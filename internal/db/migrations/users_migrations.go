package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/igosantana/igo-refrigeration-api/internal/modules/user"
	"gorm.io/gorm"
)

func getUserMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "20250418_create_users",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&user.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
	}
}

package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/igosantana/igo-refrigeration-api/internal/modules/user"
	"gorm.io/gorm"
)

func getAddressMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "20250418_create_address",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&user.Address{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("addresses")
			},
		},
	}
}

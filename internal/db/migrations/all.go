package migrations

import "github.com/go-gormigrate/gormigrate/v2"

func GetAllMigrations() []*gormigrate.Migration {
	var all []*gormigrate.Migration
	all = append(all, getUserMigrations()...)
	all = append(all, getAddressMigrations()...)
	return all
}

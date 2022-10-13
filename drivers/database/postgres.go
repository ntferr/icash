package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresInit(dsn string) gorm.Dialector {
	return postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})
}

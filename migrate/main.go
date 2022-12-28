package main

import (
	"github.com/ntferr/icash/drivers"
	"github.com/ntferr/icash/entities"
)

func main() {
	db := drivers.InitDrivers().GormDb

	db.AutoMigrate(
		&entities.Bank{},
		&entities.Card{},
		&entities.Debt{},
		&entities.Ticket{},
	)
}

package main

import (
	"log"

	"github.com/ntferr/icash/drivers"
	"github.com/ntferr/icash/entities"
)

func main() {
	db := drivers.InitDrivers().GormDb

	log.Println("init automigrate")

	err := db.AutoMigrate(
		&entities.Bank{},
		&entities.Card{},
		&entities.Ticket{},
		&entities.Debt{},
	)
	if err != nil {
		log.Printf("failed to do automigrate: %s", err.Error())
	}

	log.Println("sucessfuly automigrate!")

}

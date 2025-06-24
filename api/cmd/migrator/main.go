package main

import (
	"log"

	"github.com/icash/internal/controller/bank"
	"github.com/icash/internal/controller/card"
	"github.com/icash/internal/controller/debt"
	"github.com/icash/internal/controller/ticket"
	"github.com/icash/pkg/drivers"
)

func main() {
	db := drivers.InitDrivers().GormDB

	log.Println("init automigrate")

	err := db.AutoMigrate(
		&bank.Bank{},
		&card.Card{},
		&ticket.Ticket{},
		&debt.Debt{},
		&debt.Installment{},
	)

	if err != nil {
		log.Printf("failed to do automigrate: %s", err.Error())
	} else {
		log.Println("sucessfuly automigrate!")
	}

}

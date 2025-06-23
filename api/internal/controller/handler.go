package controller

import (
	"fmt"
	"os"

	bank_controller "github.com/ntferr/icash/controller/bank"
	card_controller "github.com/ntferr/icash/controller/card"
	debt_controller "github.com/ntferr/icash/controller/debt"
	health_controller "github.com/ntferr/icash/controller/health"
	ticket_controller "github.com/ntferr/icash/controller/ticket"
	"github.com/ntferr/icash/drivers"
	"github.com/ntferr/icash/entities"
	bank_service "github.com/ntferr/icash/service/bank"
	"github.com/ntferr/icash/service/crud"
	health_service "github.com/ntferr/icash/service/health"
)

type Controllers struct {
	Health health_controller.Controller
	Bank   bank_controller.Controller
	Card   card_controller.Controller
	Debt   debt_controller.Controller
	Ticket ticket_controller.Controller
}

func Init(drv *drivers.Drivers) Controllers {
	gorm := drv.GormDB

	sqlDb, err := gorm.DB()
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to get *sql.DB: %s", err.Error())
	}

	return Controllers{
		Health: health_controller.NewController(health_service.NewService(sqlDb)),
		Bank:   bank_controller.NewController(bank_service.NewBankCRUD(gorm)),
		Card:   card_controller.NewController(crud.NewCrud[entities.Card](gorm)),
		Debt:   debt_controller.NewController(crud.NewCrud[entities.Debt](gorm)),
		Ticket: ticket_controller.NewController(crud.NewCrud[entities.Ticket](gorm)),
	}
}

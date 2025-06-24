package controller

import (
	"fmt"
	"os"

	"github.com/icash/internal/controller/bank"
	"github.com/icash/internal/controller/card"
	"github.com/icash/internal/controller/debt"
	"github.com/icash/internal/controller/health"
	"github.com/icash/internal/controller/ticket"
	"github.com/icash/pkg/drivers"
)

type Controllers struct {
	Health health.Controller
	Bank   bank.Controller
	Card   card.Controller
	Debt   debt.Controller
	Ticket ticket.Controller
}

func Init(drv *drivers.Drivers) Controllers {
	gorm := drv.GormDB

	sqlDb, err := gorm.DB()
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to get *sql.DB: %s", err.Error())
	}

	return Controllers{
		Health: health.NewController(health_service.NewService(sqlDb)),
		Bank:   bank.NewController(bank_service.NewBankCRUD(gorm)),
		Card:   card.NewController(crud.NewCrud[entities.Card](gorm)),
		Debt:   debt.NewController(crud.NewCrud[entities.Debt](gorm)),
		Ticket: ticket.NewController(crud.NewCrud[entities.Ticket](gorm)),
	}
}

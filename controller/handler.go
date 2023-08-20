package controller

import (
	"fmt"
	"os"

	bank_controller "github.com/ntferr/icash/controller/bank"
	card_controller "github.com/ntferr/icash/controller/card"
	health_controller "github.com/ntferr/icash/controller/health"
	"github.com/ntferr/icash/drivers"
	"github.com/ntferr/icash/entities"
	"github.com/ntferr/icash/service/crud"
	health_service "github.com/ntferr/icash/service/health"
)

type Controllers struct {
	Health health_controller.Controller
	Bank   bank_controller.Controller
	Card   card_controller.Controller
}

func Init(drv *drivers.Drivers) Controllers {
	gorm := drv.GormDb

	sqlDb, err := gorm.DB()
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to get *sql.DB: %s", err.Error())
	}

	return Controllers{
		Health: health_controller.NewController(health_service.NewService(sqlDb)),
		Bank:   bank_controller.NewController(crud.NewCrud[entities.Bank](gorm)),
		Card:   card_controller.NewController(crud.NewCrud[entities.Card](gorm)),
	}
}

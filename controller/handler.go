package controller

import (
	"fmt"
	"os"

	bank_c "github.com/ntferr/icash/controller/bank"
	health_c "github.com/ntferr/icash/controller/health"
	"github.com/ntferr/icash/drivers"
	"github.com/ntferr/icash/entities"
	"github.com/ntferr/icash/service/crud"
	health_s "github.com/ntferr/icash/service/health"
)

type Controllers struct {
	Health health_c.Controller
	Bank   bank_c.Controller
}

func Init(drv *drivers.Drivers) Controllers {
	gorm := drv.GormDb

	sqlDb, err := gorm.DB()
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to get *sql.DB: %s", err.Error())
	}

	return Controllers{
		Health: health_c.NewController(health_s.NewService(sqlDb)),
		Bank:   bank_c.NewController(crud.NewCrud[entities.Bank](gorm)),
	}
}

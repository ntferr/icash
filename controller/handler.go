package controller

import (
	"fmt"
	"os"

	bank_c "github.com/ntferr/icash/controller/bank"
	health_c "github.com/ntferr/icash/controller/health"
	"github.com/ntferr/icash/drivers"
	bank_s "github.com/ntferr/icash/service/bank"
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
		Bank:   bank_c.NewController(bank_s.NewService(gorm)),
	}
}

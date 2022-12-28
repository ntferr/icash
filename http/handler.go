package http

import (
	"fmt"
	"os"

	healthc "github.com/ntferr/icash/controller/health"
	"github.com/ntferr/icash/drivers"
	healths "github.com/ntferr/icash/service/health"
)

type Controllers struct {
	health healthc.Controller
}

func InitControllers(drv *drivers.Drivers) Controllers {
	sqlDb, err := drv.GormDb.DB()
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to get *sql.DB: %s", err.Error())
	}

	return Controllers{
		health: healthc.NewController(healths.NewService(sqlDb)),
	}
}

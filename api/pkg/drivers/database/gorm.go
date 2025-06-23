package database

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

func GormInit(dialector gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(dialector)
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to open connection: %s", err)
	}

	return db
}

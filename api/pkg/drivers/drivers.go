package drivers

import (
	"fmt"

	"github.com/icash/internal/settings"
	"github.com/icash/pkg/drivers/database"
	"gorm.io/gorm"
)

type Drivers struct {
	GormDB *gorm.DB
}

// InitDrivers: init drivers
func InitDrivers() Drivers {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		settings.GetSettings().Postgres.Host,
		settings.GetSettings().Postgres.User,
		settings.GetSettings().Postgres.Password,
		settings.GetSettings().Postgres.Name,
		settings.GetSettings().Postgres.Port,
	)

	gormDialector := database.PostgresInit(dsn)
	gormDB := database.GormInit(gormDialector)

	return Drivers{
		GormDB: gormDB,
	}
}

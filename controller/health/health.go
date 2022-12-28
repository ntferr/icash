package health

import (
	"github.com/gofiber/fiber/v2"
	healths "github.com/ntferr/icash/service/health"
)

type health struct {
	service healths.Service
}

type Controller interface {
	Status(c *fiber.Ctx) error
}

func NewController(service healths.Service) Controller {
	return &health{
		service: service,
	}
}

func (h health) Status(c *fiber.Ctx) error {
	var dbStr string

	DBStatus := h.service.CheckDatabase()
	if DBStatus {
		dbStr = "Database is healthy"
	} else {
		dbStr = "Database is unhealthy"
	}

	return c.JSON(fiber.Map{
		"Database": dbStr,
		"Cache":    "",
		"Service":  validate(DBStatus),
	})
}

func validate(db bool) string {
	// TODO add a validate to cache too
	if db {
		return "Service is healthy"
	}

	return "Service is unhealthy"
}

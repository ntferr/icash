package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ntferr/icash/controller"
)

func SetupRouter(api fiber.Router, c controller.Controllers) {
	api.Get("/health-check", c.Health.Status)

	//Bank endpoints
	bank_router := api.Group("/banks")
	bank_router.Get("/all", c.Bank.FindAll)
	bank_router.Get("/:id", c.Bank.FindByID)
	bank_router.Post("", c.Bank.New)
	bank_router.Put("/:id", c.Bank.Alter)
	bank_router.Delete("/:id", c.Bank.Remove)
}
